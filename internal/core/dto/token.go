package dto

import (
	"time"

	"github.com/google/uuid"
)

type TokenPayload struct {
	ID        uuid.UUID `json:"id"`
	IssuedAt  time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expiry_at"`
}

func NewPayload(id uuid.UUID, duration time.Duration) *TokenPayload {
	payload := &TokenPayload{
		ID:        id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload
}

func (p *TokenPayload) IsNotExpired() bool {
	return time.Now().Before(p.ExpiredAt)
}
