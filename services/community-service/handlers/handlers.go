package handlers

import (
	"community-service/config"
	"community-service/models"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "community-service", "time": time.Now()})
}

func GetCommunities(c *gin.Context) {
	userID := c.GetString("user_id")

	var communities []models.Community
	config.DB.Table("communities").
		Joins("JOIN community_members ON communities.id = community_members.community_id").
		Where("community_members.user_id = ?", userID).
		Find(&communities)

	c.JSON(http.StatusOK, gin.H{"communities": communities})
}

func CreateCommunity(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateCommunityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userUUID, _ := uuid.Parse(userID)
	community := models.Community{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   userUUID,
		InviteCode:  fmt.Sprintf("%06d", rand.Intn(1000000)),
	}

	config.DB.Create(&community)

	// Add creator as admin
	member := models.CommunityMember{
		CommunityID: community.ID,
		UserID:      userUUID,
		Role:        "admin",
		JoinedAt:    time.Now(),
	}
	config.DB.Create(&member)

	c.JSON(http.StatusCreated, gin.H{"message": "Community created", "community": community})
}

func JoinCommunity(c *gin.Context) {
	userID := c.GetString("user_id")

	var req struct {
		InviteCode string `json:"invite_code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var community models.Community
	if err := config.DB.Where("invite_code = ?", req.InviteCode).First(&community).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid invite code"})
		return
	}

	userUUID, _ := uuid.Parse(userID)
	member := models.CommunityMember{
		CommunityID: community.ID,
		UserID:      userUUID,
		Role:        "member",
		JoinedAt:    time.Now(),
	}

	config.DB.Create(&member)
	c.JSON(http.StatusOK, gin.H{"message": "Joined community successfully", "community": community})
}
