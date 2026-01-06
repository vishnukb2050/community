package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title     string    `gorm:"size:255" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	NoteType  string    `gorm:"size:50;default:'note'" json:"note_type"` // note, checklist
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChecklistItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NoteID    uuid.UUID `gorm:"type:uuid;not null" json:"note_id"`
	ItemText  string    `gorm:"size:255;not null" json:"item_text"`
	IsChecked bool      `gorm:"default:false" json:"is_checked"`
	Position  int       `gorm:"default:0" json:"position"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateNoteRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	NoteType string `json:"note_type"` // note, checklist
}

type AddChecklistItemRequest struct {
	ItemText string `json:"item_text" binding:"required"`
}
