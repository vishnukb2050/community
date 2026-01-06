package models

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Type          string    `gorm:"size:50" json:"type"` // private, community
	CommunityID   uuid.UUID `gorm:"type:uuid" json:"community_id,omitempty"`
	LastMessageAt time.Time `json:"last_message_at"`
	CreatedAt     time.Time `json:"created_at"`
}

type Message struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	ConversationID uuid.UUID `gorm:"type:uuid;not null" json:"conversation_id"`
	SenderID       uuid.UUID `gorm:"type:uuid;not null" json:"sender_id"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	SentAt         time.Time `json:"sent_at"`
}

type ConversationParticipant struct {
	ConversationID uuid.UUID `gorm:"type:uuid;primary_key" json:"conversation_id"`
	UserID         uuid.UUID `gorm:"type:uuid;primary_key" json:"user_id"`
	JoinedAt       time.Time `json:"joined_at"`
}

type CreateConversationRequest struct {
	ParticipantID string `json:"participant_id" binding:"required"`
}

type SendMessageRequest struct {
	Content string `json:"content" binding:"required"`
}
