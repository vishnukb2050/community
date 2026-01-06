package models

import (
	"time"

	"github.com/google/uuid"
)

type Reminder struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title          string    `gorm:"size:255;not null" json:"title"`
	Amount         float64   `json:"amount"`
	DueDate        time.Time `gorm:"not null" json:"due_date"`
	IsRecurring    bool      `gorm:"default:false" json:"is_recurring"`
	RecurrenceType string    `gorm:"size:50" json:"recurrence_type"`          // monthly, yearly, custom
	Status         string    `gorm:"size:50;default:'pending'" json:"status"` // pending, paid, snoozed, overdue
	SnoozedUntil   time.Time `json:"snoozed_until,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateReminderRequest struct {
	Title          string  `json:"title" binding:"required"`
	Amount         float64 `json:"amount"`
	DueDate        string  `json:"due_date" binding:"required"` // YYYY-MM-DD
	IsRecurring    bool    `json:"is_recurring"`
	RecurrenceType string  `json:"recurrence_type"` // monthly, yearly
}

type UpdateReminderRequest struct {
	Title   string  `json:"title"`
	Amount  float64 `json:"amount"`
	DueDate string  `json:"due_date"`
	Status  string  `json:"status"`
}
