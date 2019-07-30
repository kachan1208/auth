package model

import (
	"time"
)

type Token struct {
	ID        string    `json:"id"`
	AccountID string    `json:"-"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"-"`
}

func (t *Token) IsDeleted() bool {
	return !(t.DeletedAt == nilTime)
}
