package models

import (
	"time"

	"github.com/google/uuid"
)

type Poll struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CommunityID uuid.UUID `gorm:"type:uuid;not null" json:"community_id"`
	Question    string    `gorm:"size:500;not null" json:"question"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	EndsAt      time.Time `json:"ends_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type PollOption struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	PollID     uuid.UUID `gorm:"type:uuid;not null" json:"poll_id"`
	OptionText string    `gorm:"size:255;not null" json:"option_text"`
	VoteCount  int       `gorm:"default:0" json:"vote_count"`
}

type PollVote struct {
	PollID   uuid.UUID `gorm:"type:uuid;primary_key" json:"poll_id"`
	UserID   uuid.UUID `gorm:"type:uuid;primary_key" json:"user_id"`
	OptionID uuid.UUID `gorm:"type:uuid;not null" json:"option_id"`
	VotedAt  time.Time `json:"voted_at"`
}

type CreatePollRequest struct {
	CommunityID string   `json:"community_id" binding:"required"`
	Question    string   `json:"question" binding:"required"`
	Options     []string `json:"options" binding:"required,min=2"`
}
