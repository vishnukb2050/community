package handlers

import (
	"net/http"
	"notice-service/config"
	"notice-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "notice-service", "time": time.Now()})
}

func GetNotices(c *gin.Context) {
	communityID := c.Query("community_id")

	if communityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "community_id required"})
		return
	}

	var notices []models.Notice
	config.DB.Where("community_id = ?", communityID).Order("created_at DESC").Find(&notices)

	c.JSON(http.StatusOK, gin.H{"notices": notices, "count": len(notices)})
}

func CreateNotice(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateNoticeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	communityUUID, _ := uuid.Parse(req.CommunityID)
	userUUID, _ := uuid.Parse(userID)

	notice := models.Notice{
		CommunityID: communityUUID,
		Title:       req.Title,
		Content:     req.Content,
		PostedBy:    userUUID,
	}

	if err := config.DB.Create(&notice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notice"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Notice posted successfully", "notice": notice})
}

func GetNotice(c *gin.Context) {
	noticeID := c.Param("id")

	var notice models.Notice
	if err := config.DB.Where("id = ?", noticeID).First(&notice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notice not found"})
		return
	}

	c.JSON(http.StatusOK, notice)
}

func DeleteNotice(c *gin.Context) {
	userID := c.GetString("user_id")
	noticeID := c.Param("id")

	// Check if user posted the notice
	result := config.DB.Where("id = ? AND posted_by = ?", noticeID, userID).Delete(&models.Notice{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this notice"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notice deleted successfully"})
}
