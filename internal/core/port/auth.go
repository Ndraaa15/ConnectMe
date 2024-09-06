package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/google/uuid"
)

type AuthRepositoryItf interface {
	NewAuthRepositoryClient(tx bool) AuthRepositoryClientItf
}

type AuthRepositoryClientItf interface {
	Commit() error
	Rollback() error
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
}

type AuthServiceItf interface {
	Register(ctx context.Context, req dto.SignUpRequest) (uuid.UUID, error)
	Verify(ctx context.Context, req dto.VerifyAccountRequest) (uuid.UUID, error)
}
