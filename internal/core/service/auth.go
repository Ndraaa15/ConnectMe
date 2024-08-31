package service

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/google/uuid"
)

type AuthService struct {
	repository port.AuthRepositoryItf
	cache      port.RedisItf
}

func NewAuthService(repository port.AuthRepositoryItf, cache port.RedisItf) *AuthService {
	return &AuthService{
		repository: repository,
		cache:      cache,
	}
}

func (auth *AuthService) Register(ctx context.Context, req dto.SignUpRequest) (uuid.UUID, error) {
	return uuid.New(), nil
}
