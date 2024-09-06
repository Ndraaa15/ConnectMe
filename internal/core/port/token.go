package port

import "github.com/Ndraaa15/ConnectMe/internal/core/dto"

type TokenItf interface {
	Decode(token string) (*dto.TokenPayload, error)
	Encode(payload dto.TokenPayload) (string, error)
}
