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

func GetCommunity(c *gin.Context) {
	communityID := c.Param("id")

	var community models.Community
	if err := config.DB.Where("id = ?", communityID).First(&community).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Community not found"})
		return
	}

	var members []struct {
		models.CommunityMember
		Name   string `json:"name"`
		Mobile string `json:"mobile"`
	}

	config.DB.Table("community_members").
		Select("community_members.*, users.name, users.mobile").
		Joins("JOIN users ON community_members.user_id = users.id").
		Where("community_members.community_id = ?", communityID).
		Find(&members)

	c.JSON(http.StatusOK, gin.H{
		"community": community,
		"members":   members,
	})
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

func AddMember(c *gin.Context) {
	userID := c.GetString("user_id")

	var req struct {
		CommunityID string `json:"community_id"`
		Mobile      string `json:"mobile"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify Admin
	var adminMember models.CommunityMember
	if err := config.DB.Where("community_id = ? AND user_id = ? AND role = 'admin'", req.CommunityID, userID).First(&adminMember).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can add members"})
		return
	}

	// Find Target User
	var targetUser models.User
	if err := config.DB.Table("users").Where("mobile = ?", req.Mobile).First(&targetUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found with this mobile number"})
		return
	}

	// Check if already member
	var existingMember models.CommunityMember
	if err := config.DB.Where("community_id = ? AND user_id = ?", req.CommunityID, targetUser.ID).First(&existingMember).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already a member"})
		return
	}

	// Add Member
	newMember := models.CommunityMember{
		CommunityID: uuid.MustParse(req.CommunityID),
		UserID:      targetUser.ID,
		Role:        "member",
		JoinedAt:    time.Now(),
	}
	config.DB.Create(&newMember)

	c.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}
