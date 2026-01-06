package models

import (
	"time"

	"github.com/google/uuid"
)

type Notice struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CommunityID uuid.UUID `gorm:"type:uuid;not null" json:"community_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	PostedBy    uuid.UUID `gorm:"type:uuid;not null" json:"posted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateNoticeRequest struct {
	CommunityID string `json:"community_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
}
