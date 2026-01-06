package models

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID          uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Amount          float64   `gorm:"not null" json:"amount"`
	Category        string    `gorm:"size:100" json:"category"`
	Description     string    `json:"description"`
	Date            time.Time `gorm:"not null" json:"date"`
	ReceiptImageURL string    `json:"receipt_image_url,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateExpenseRequest struct {
	Amount      float64 `json:"amount" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description"`
	Date        string  `json:"date" binding:"required"` // YYYY-MM-DD
}
