package postgresql

import (
	"context"
	"errors"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) NewAuthRepositoryClient(tx bool) port.AuthRepositoryClientItf {
	if tx {
		return &AuthRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &AuthRepositoryClient{
			q: r.db,
		}
	}
}

type AuthRepositoryClient struct {
	q *gorm.DB
}

func (r *AuthRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *AuthRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *AuthRepositoryClient) CreateUser(ctx context.Context, user *domain.User) error {
	return r.q.Debug().WithContext(ctx).Model(&domain.User{}).Create(user).Error
}

func (r *AuthRepositoryClient) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.q.Debug().WithContext(ctx).Model(&domain.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &domain.User{}, errx.New(fiber.StatusNotFound, "user not found", err)
		}
	}

	return &user, nil
}

func (r *AuthRepositoryClient) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.q.Debug().WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &domain.User{}, errx.New(fiber.StatusNotFound, "user not found", err)
		}
		return &domain.User{}, err
	}

	return &user, nil
}

func (r *AuthRepositoryClient) UpdateUser(ctx context.Context, user *domain.User) error {
	return r.q.Debug().WithContext(ctx).Model(&domain.User{}).Where("id = ?", user.ID).Save(user).Error
}
