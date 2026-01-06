package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Mobile            string    `gorm:"uniqueIndex;size:15" json:"mobile"`
	Name              string    `gorm:"size:255" json:"name"`
	Email             string    `gorm:"size:255" json:"email"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	AllowChatRequests bool      `json:"allow_chat_requests"`
	HidePhoneNumber   bool      `json:"hide_phone_number"`
	Budget            float64   `gorm:"default:0.00" json:"budget"`
}

type UpdateProfileRequest struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Budget float64 `json:"budget"`
}

type UpdateSettingsRequest struct {
	AllowChatRequests bool `json:"allow_chat_requests"`
	HidePhoneNumber   bool `json:"hide_phone_number"`
}
