package handlers

import (
	"net/http"
	"notification-service/config"
	"notification-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "notification-service",
		"time":    time.Now(),
	})
}

// GetNotifications - List user notifications
func GetNotifications(c *gin.Context) {
	userID := c.GetString("user_id")
	unreadOnly := c.Query("unread") == "true"

	query := config.DB.Where("user_id = ?", userID)
	if unreadOnly {
		query = query.Where("is_read = false")
	}

	var notifications []models.Notification
	if err := query.Order("sent_at DESC").
		Limit(50).
		Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifications, "count": len(notifications)})
}

// SendNotification - Send notification to user
func SendNotification(c *gin.Context) {
	var req models.SendNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userUUID, _ := uuid.Parse(req.UserID)
	notification := models.Notification{
		UserID: userUUID,
		Title:  req.Title,
		Body:   req.Body,
		Type:   req.Type,
		SentAt: time.Now(),
	}

	if err := config.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	// In production: Send push notification via Firebase
	// sendPushNotification(req.UserID, req.Title, req.Body)

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Notification sent successfully",
		"notification": notification,
	})
}

// MarkAsRead - Mark notification as read
func MarkAsRead(c *gin.Context) {
	userID := c.GetString("user_id")
	notifID := c.Param("id")

	var notification models.Notification
	if err := config.DB.Where("id = ? AND user_id = ?", notifID, userID).
		First(&notification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	notification.IsRead = true
	notification.ReadAt = time.Now()
	config.DB.Save(&notification)

	c.JSON(http.StatusOK, gin.H{"message": "Marked as read"})
}

// MarkAllAsRead - Mark all notifications as read
func MarkAllAsRead(c *gin.Context) {
	userID := c.GetString("user_id")

	config.DB.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = false", userID).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": time.Now(),
		})

	c.JSON(http.StatusOK, gin.H{"message": "All notifications marked as read"})
}

// RegisterDevice - Register device token for push notifications
func RegisterDevice(c *gin.Context) {
	userID := c.GetString("user_id")

	var req struct {
		DeviceToken string `json:"device_token" binding:"required"`
		Platform    string `json:"platform" binding:"required"` // android, ios
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userUUID, _ := uuid.Parse(userID)
	deviceToken := models.DeviceToken{
		UserID:      userUUID,
		DeviceToken: req.DeviceToken,
		Platform:    req.Platform,
		UpdatedAt:   time.Now(),
	}

	// Upsert device token
	config.DB.Save(&deviceToken)

	c.JSON(http.StatusOK, gin.H{"message": "Device registered successfully"})
}

// DeleteNotification - Delete notification
func DeleteNotification(c *gin.Context) {
	userID := c.GetString("user_id")
	notifID := c.Param("id")

	result := config.DB.Where("id = ? AND user_id = ?", notifID, userID).
		Delete(&models.Notification{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted"})
}

// Helper function (to be implemented with Firebase SDK)
func sendPushNotification(userID, title, body string) error {
	// Get device token
	var deviceToken models.DeviceToken
	if err := config.DB.Where("user_id = ?", userID).First(&deviceToken).Error; err != nil {
		return err
	}

	// Send via Firebase Cloud Messaging
	// This requires Firebase Admin SDK integration
	// firebase.SendMessage(deviceToken.DeviceToken, title, body)

	return nil
}
