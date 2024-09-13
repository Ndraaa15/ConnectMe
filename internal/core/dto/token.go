package dto

import (
	"time"

	"github.com/google/uuid"
)

type TokenPayload struct {
	ID        uuid.UUID `json:"id"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
}

func NewPayload(id uuid.UUID, duration time.Duration) TokenPayload {
	payload := TokenPayload{
		ID:        id,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload
}

func (p *TokenPayload) IsNotExpired() bool {
	return time.Now().Before(p.ExpiresAt)
}
