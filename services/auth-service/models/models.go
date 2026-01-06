package models

import (
	"time"

	"github.com/google/uuid"
)

// User model
type User struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Mobile            string    `gorm:"uniqueIndex;size:15;not null" json:"mobile"`
	Name              string    `gorm:"size:255" json:"name"`
	Email             string    `gorm:"size:255" json:"email"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	AllowChatRequests bool      `gorm:"default:true" json:"allow_chat_requests"`
	HidePhoneNumber   bool      `gorm:"default:false" json:"hide_phone_number"`
}

// OTPCode model
type OTPCode struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Mobile    string    `gorm:"size:15;not null" json:"mobile"`
	OTPCode   string    `gorm:"size:6;not null" json:"otp_code"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Verified  bool      `gorm:"default:false" json:"verified"`
	CreatedAt time.Time `json:"created_at"`
}

// Request/Response models
type SendOTPRequest struct {
	Mobile string `json:"mobile" binding:"required"`
}

type VerifyOTPRequest struct {
	Mobile string `json:"mobile" binding:"required"`
	OTP    string `json:"otp" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
