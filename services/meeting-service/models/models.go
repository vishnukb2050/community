package models

import (
	"time"

	"github.com/google/uuid"
)

type MeetingMinute struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CommunityID uuid.UUID `gorm:"type:uuid;not null" json:"community_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	MeetingDate time.Time `gorm:"not null" json:"meeting_date"`
	UploadedBy  uuid.UUID `gorm:"type:uuid;not null" json:"uploaded_by"`
	FileURL     string    `json:"file_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateMinuteRequest struct {
	CommunityID string `json:"community_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	MeetingDate string `json:"meeting_date" binding:"required"` // YYYY-MM-DD
}
