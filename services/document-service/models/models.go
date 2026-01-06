package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	FileName    string    `gorm:"size:255;not null" json:"file_name"`
	FileURL     string    `gorm:"size:500;not null" json:"file_url"`
	FileType    string    `gorm:"size:50" json:"file_type"` // pdf, image, doc
	FileSize    int64     `json:"file_size"`                // in bytes
	Category    string    `gorm:"size:100" json:"category,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type UploadDocumentRequest struct {
	FileName    string `json:"file_name" binding:"required"`
	FileBase64  string `json:"file_base64" binding:"required"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
