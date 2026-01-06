package handlers

import (
	"meeting-service/config"
	"meeting-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "meeting-service", "time": time.Now()})
}

func GetMinutes(c *gin.Context) {
	communityID := c.Query("community_id")

	if communityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "community_id required"})
		return
	}

	var minutes []models.MeetingMinute
	config.DB.Where("community_id = ?", communityID).Order("meeting_date DESC").Find(&minutes)

	c.JSON(http.StatusOK, gin.H{"minutes": minutes, "count": len(minutes)})
}

func CreateMinute(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateMinuteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meetingDate, err := time.Parse("2006-01-02", req.MeetingDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	communityUUID, _ := uuid.Parse(req.CommunityID)
	userUUID, _ := uuid.Parse(userID)

	minute := models.MeetingMinute{
		CommunityID: communityUUID,
		Title:       req.Title,
		Content:     req.Content,
		MeetingDate: meetingDate,
		UploadedBy:  userUUID,
	}

	if err := config.DB.Create(&minute).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create minute"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Meeting minute created successfully", "minute": minute})
}

func GetMinute(c *gin.Context) {
	minuteID := c.Param("id")

	var minute models.MeetingMinute
	if err := config.DB.Where("id = ?", minuteID).First(&minute).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Minute not found"})
		return
	}

	c.JSON(http.StatusOK, minute)
}

func DeleteMinute(c *gin.Context) {
	userID := c.GetString("user_id")
	minuteID := c.Param("id")

	result := config.DB.Where("id = ? AND uploaded_by = ?", minuteID, userID).Delete(&models.MeetingMinute{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meeting minute deleted successfully"})
}
