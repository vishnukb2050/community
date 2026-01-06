package handlers

import (
	"net/http"
	"profile-service/config"
	"profile-service/models"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "profile-service",
		"time":    time.Now(),
	})
}

func GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user":    user,
	})
}

func UpdateSettings(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.UpdateSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.AllowChatRequests = req.AllowChatRequests
	user.HidePhoneNumber = req.HidePhoneNumber

	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Settings updated successfully",
	})
}

func GetStats(c *gin.Context) {
	userID := c.GetString("user_id")

	var expenseCount int64
	var reminderCount int64

	config.DB.Table("expenses").Where("user_id = ?", userID).Count(&expenseCount)
	config.DB.Table("reminders").Where("user_id = ?", userID).Count(&reminderCount)

	c.JSON(http.StatusOK, gin.H{
		"total_expenses":  expenseCount,
		"total_reminders": reminderCount,
	})
}
