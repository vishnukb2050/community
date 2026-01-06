package models

import (
	"time"

	"github.com/google/uuid"
)

type Community struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `json:"description"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	InviteCode  string    `gorm:"size:50;unique" json:"invite_code"`
	CreatedAt   time.Time `json:"created_at"`
}

type CommunityMember struct {
	CommunityID uuid.UUID `gorm:"type:uuid;primary_key" json:"community_id"`
	UserID      uuid.UUID `gorm:"type:uuid;primary_key" json:"user_id"`
	Role        string    `gorm:"size:50;default:'member'" json:"role"` // admin, member
	JoinedAt    time.Time `json:"joined_at"`
}

type CreateCommunityRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type User struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Mobile string    `json:"mobile"`
}
