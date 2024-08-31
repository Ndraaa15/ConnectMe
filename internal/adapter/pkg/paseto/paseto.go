package paseto

import (
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	IssuedAt  time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expiry_at"`
}

func NewPayload(id uuid.UUID, duration time.Duration) *Payload {
	payload := &Payload{
		ID:        id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload
}

func (p *Payload) IsNotExpired() bool {
	return time.Now().Before(p.ExpiredAt)
}

type PasetoMaker struct {
	key          paseto.V4SymmetricKey
	pasetoParser paseto.Parser
}

func NewPaseto() *PasetoMaker {
	key := paseto.NewV4SymmetricKey()

	pasetoParser := paseto.NewParser()
	pasetoParser.AddRule(paseto.IssuedBy("ConnectMe"))
	pasetoParser.AddRule(paseto.Subject("Authentication"))

	return &PasetoMaker{
		key:          key,
		pasetoParser: pasetoParser,
	}
}

func (p *PasetoMaker) Encode(payload Payload) (string, error) {
	token := paseto.NewToken()
	token.Set("payload", payload)

	tokenEncrypted := token.V4Encrypt(p.key, nil)

	return tokenEncrypted, nil

}

func (p *PasetoMaker) Decode(token string) (*Payload, error) {
	tok, err := p.pasetoParser.ParseV4Local(p.key, token, nil)
	if err != nil {
		return &Payload{}, err
	}

	var payload *Payload
	err = tok.Get("payload", payload)
	if err != nil {
		return &Payload{}, err
	}

	if !payload.IsNotExpired() {
		return &Payload{}, errx.New(fiber.StatusUnauthorized, "Token is invalid", errors.New("TOKEN_EXPIRED"))
	}

	return payload, nil
}
