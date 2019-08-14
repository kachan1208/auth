package model

import (
	"time"
)

type Token struct {
	ID        string    `json:"id"`
	AccountID string    `json:"-"`
	IsEnabled bool      `json:"is_enabled"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"-"`
}

func (t *Token) IsDeleted() bool {
	return !(t.DeletedAt == nilTime)
}
