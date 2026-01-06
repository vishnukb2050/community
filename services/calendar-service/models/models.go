package models

import (
	"time"

	"github.com/google/uuid"
)

type CalendarEvent struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `json:"description"`
	EventDate   time.Time `gorm:"not null" json:"event_date"`
	EventTime   time.Time `json:"event_time"`
	EventType   string    `gorm:"size:50;default:'personal'" json:"event_type"` // personal, reminder, community
	RelatedID   uuid.UUID `gorm:"type:uuid" json:"related_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateEventRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	EventDate   string `json:"event_date" binding:"required"` // YYYY-MM-DD format
	EventTime   string `json:"event_time"`                    // HH:MM format
}
