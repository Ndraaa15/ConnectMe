package paseto

import (
	"errors"

	"aidanwoods.dev/go-paseto"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Paseto struct {
	key          paseto.V4SymmetricKey
	pasetoParser paseto.Parser
}

func NewPaseto(env env.Token) *Paseto {
	key, err := paseto.V4SymmetricKeyFromHex(env.Secret)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create paseto key")
	}

	pasetoParser := paseto.NewParser()
	pasetoParser.AddRule(paseto.IssuedBy("ConnectMe"))
	pasetoParser.AddRule(paseto.Subject("Authentication"))

	return &Paseto{
		key:          key,
		pasetoParser: pasetoParser,
	}
}

func (p *Paseto) Encode(payload dto.TokenPayload) (string, error) {
	token := paseto.NewToken()
	err := token.Set("payload", payload)
	if err != nil {
		return "", err
	}
	token.SetIssuedAt(payload.IssuedAt)
	token.SetExpiration(payload.ExpiresAt)
	token.SetIssuer("ConnectMe")
	token.SetSubject("Authentication")

	tokenEncrypted := token.V4Encrypt(p.key, nil)

	return tokenEncrypted, nil

}

func (p *Paseto) Decode(token string) (dto.TokenPayload, error) {
	tok, err := p.pasetoParser.ParseV4Local(p.key, token, nil)
	if err != nil {
		return dto.TokenPayload{}, err
	}

	var payload dto.TokenPayload
	err = tok.Get("payload", &payload)
	if err != nil {
		return dto.TokenPayload{}, err
	}

	if !payload.IsNotExpired() {
		return dto.TokenPayload{}, errx.New(fiber.StatusUnauthorized, "token is invalid", errors.New("TOKEN_EXPIRED"))
	}

	return payload, nil
}
