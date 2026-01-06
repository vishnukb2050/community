package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Body      string    `gorm:"type:text;not null" json:"body"`
	Type      string    `gorm:"size:50" json:"type"` // reminder, notice, poll, chat, system
	RelatedID uuid.UUID `gorm:"type:uuid" json:"related_id,omitempty"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	SentAt    time.Time `json:"sent_at"`
	ReadAt    time.Time `json:"read_at,omitempty"`
}

type SendNotificationRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Type   string `json:"type"`
}

type DeviceToken struct {
	UserID      uuid.UUID `gorm:"type:uuid;primary_key" json:"user_id"`
	DeviceToken string    `gorm:"size:500;not null" json:"device_token"`
	Platform    string    `gorm:"size:20" json:"platform"` // android, ios
	UpdatedAt   time.Time `json:"updated_at"`
}
