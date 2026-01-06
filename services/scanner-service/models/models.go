package models

import (
	"time"

	"github.com/google/uuid"
)

type ScannedBill struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	ImageURL       string    `gorm:"size:500" json:"image_url"`
	ExtractedText  string    `gorm:"type:text" json:"extracted_text"`
	DetectedAmount float64   `json:"detected_amount"`
	DetectedDate   time.Time `json:"detected_date,omitempty"`
	DetectedVendor string    `gorm:"size:255" json:"detected_vendor,omitempty"`
	Category       string    `gorm:"size:100" json:"category,omitempty"`
	IsProcessed    bool      `gorm:"default:false" json:"is_processed"`
	CreatedAt      time.Time `json:"created_at"`
}

type ScanBillRequest struct {
	ImageBase64 string `json:"image_base64" binding:"required"`
}
