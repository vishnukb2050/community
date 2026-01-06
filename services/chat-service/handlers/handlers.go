package handlers

import (
	"chat-service/config"
	"chat-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "chat-service",
		"time":    time.Now(),
	})
}

// GetConversations - List all user conversations
func GetConversations(c *gin.Context) {
	userID := c.GetString("user_id")

	var conversations []models.Conversation
	if err := config.DB.Table("conversations").
		Joins("JOIN conversation_participants ON conversations.id = conversation_participants.conversation_id").
		Where("conversation_participants.user_id = ?", userID).
		Order("last_message_at DESC").
		Find(&conversations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch conversations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"conversations": conversations})
}

// CreateConversation - Start new private conversation
func CreateConversation(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if conversation already exists
	var existing models.Conversation
	err := config.DB.Table("conversations").
		Joins("JOIN conversation_participants p1 ON conversations.id = p1.conversation_id").
		Joins("JOIN conversation_participants p2 ON conversations.id = p2.conversation_id").
		Where("p1.user_id = ? AND p2.user_id = ? AND conversations.type = 'private'", userID, req.ParticipantID).
		First(&existing).Error

	if err == nil {
		// Conversation already exists
		c.JSON(http.StatusOK, gin.H{
			"message":      "Conversation already exists",
			"conversation": existing,
		})
		return
	}

	// Create new conversation
	conversation := models.Conversation{
		Type:          "private",
		LastMessageAt: time.Now(),
	}

	if err := config.DB.Create(&conversation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create conversation"})
		return
	}

	// Add participants
	userUUID, _ := uuid.Parse(userID)
	participantUUID, _ := uuid.Parse(req.ParticipantID)

	participants := []models.ConversationParticipant{
		{ConversationID: conversation.ID, UserID: userUUID, JoinedAt: time.Now()},
		{ConversationID: conversation.ID, UserID: participantUUID, JoinedAt: time.Now()},
	}

	config.DB.Create(&participants)

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Conversation created",
		"conversation": conversation,
	})
}

// GetConversation - Get conversation details
func GetConversation(c *gin.Context) {
	userID := c.GetString("user_id")
	conversationID := c.Param("id")

	// Verify user is participant
	var participant models.ConversationParticipant
	if err := config.DB.Where("conversation_id = ? AND user_id = ?", conversationID, userID).
		First(&participant).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var conversation models.Conversation
	if err := config.DB.Where("id = ?", conversationID).First(&conversation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
		return
	}

	c.JSON(http.StatusOK, conversation)
}

// GetMessages - Get conversation messages
func GetMessages(c *gin.Context) {
	userID := c.GetString("user_id")
	conversationID := c.Param("id")

	// Verify user is participant
	var participant models.ConversationParticipant
	if err := config.DB.Where("conversation_id = ? AND user_id = ?", conversationID, userID).
		First(&participant).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var messages []models.Message
	if err := config.DB.Where("conversation_id = ?", conversationID).
		Order("sent_at ASC").
		Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

// SendMessage - Send new message
func SendMessage(c *gin.Context) {
	userID := c.GetString("user_id")
	conversationID := c.Param("id")

	// Verify user is participant
	var participant models.ConversationParticipant
	if err := config.DB.Where("conversation_id = ? AND user_id = ?", conversationID, userID).
		First(&participant).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var req models.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	senderUUID, _ := uuid.Parse(userID)
	convUUID, _ := uuid.Parse(conversationID)

	message := models.Message{
		ConversationID: convUUID,
		SenderID:       senderUUID,
		Content:        req.Content,
		SentAt:         time.Now(),
	}

	if err := config.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	// Update conversation last_message_at
	config.DB.Model(&models.Conversation{}).
		Where("id = ?", conversationID).
		Update("last_message_at", time.Now())

	c.JSON(http.StatusCreated, gin.H{
		"message": "Message sent",
		"data":    message,
	})
}
