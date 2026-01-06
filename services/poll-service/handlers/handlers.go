package handlers

import (
	"net/http"
	"poll-service/config"
	"poll-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "poll-service", "time": time.Now()})
}

func GetPolls(c *gin.Context) {
	communityID := c.Query("community_id")

	if communityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "community_id required"})
		return
	}

	var polls []models.Poll
	config.DB.Where("community_id = ?", communityID).Order("created_at DESC").Find(&polls)

	c.JSON(http.StatusOK, gin.H{"polls": polls, "count": len(polls)})
}

func CreatePoll(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreatePollRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	communityUUID, _ := uuid.Parse(req.CommunityID)
	userUUID, _ := uuid.Parse(userID)

	poll := models.Poll{
		CommunityID: communityUUID,
		Question:    req.Question,
		CreatedBy:   userUUID,
	}

	if err := config.DB.Create(&poll).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create poll"})
		return
	}

	// Create options
	for _, optionText := range req.Options {
		option := models.PollOption{
			PollID:     poll.ID,
			OptionText: optionText,
			VoteCount:  0,
		}
		config.DB.Create(&option)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Poll created successfully", "poll": poll})
}

func GetPollOptions(c *gin.Context) {
	pollID := c.Param("id")

	var options []models.PollOption
	if err := config.DB.Where("poll_id = ?", pollID).Find(&options).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch options"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"options": options})
}

func Vote(c *gin.Context) {
	userID := c.GetString("user_id")
	pollID := c.Param("id")

	var req struct {
		OptionID string `json:"option_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already voted
	pollUUID, _ := uuid.Parse(pollID)
	userUUID, _ := uuid.Parse(userID)
	optionUUID, _ := uuid.Parse(req.OptionID)

	var existingVote models.PollVote
	if err := config.DB.Where("poll_id = ? AND user_id = ?", pollID, userID).First(&existingVote).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Already voted"})
		return
	}

	// Record vote
	vote := models.PollVote{
		PollID:   pollUUID,
		UserID:   userUUID,
		OptionID: optionUUID,
		VotedAt:  time.Now(),
	}

	if err := config.DB.Create(&vote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record vote"})
		return
	}

	// Increment vote count
	config.DB.Model(&models.PollOption{}).
		Where("id = ?", req.OptionID).
		UpdateColumn("vote_count", config.DB.Raw("vote_count + 1"))

	c.JSON(http.StatusOK, gin.H{"message": "Vote recorded successfully"})
}
