package dto

import (
	"time"
)

type TokenPayload struct {
	ID        string    `json:"id"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
}

func NewPayload(id string, duration time.Duration, role string) TokenPayload {
	payload := TokenPayload{
		ID:        id,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload
}

func (p *TokenPayload) IsNotExpired() bool {
	return time.Now().Before(p.ExpiresAt)
}
