package model

import (
	"time"
)

type Token struct {
	ID        string
	AccountID string
	Token     string
	CreatedAt time.Time
	DeletedAt time.Time
}

func (t *Token) IsDeleted() bool {
	return !(t.DeletedAt == nilTime)
}
