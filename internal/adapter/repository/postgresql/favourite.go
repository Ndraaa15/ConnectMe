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

func NewFavouriteRepository(db *gorm.DB) *FavouriteRepository {
	return &FavouriteRepository{
		db: db,
	}
}

type FavouriteRepository struct {
	db *gorm.DB
}

func (r *FavouriteRepository) NewFavouriteRepositoryClient(tx bool) port.FavouriteRepositoryClientItf {
	if tx {
		return &FavouriteRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &FavouriteRepositoryClient{
			q: r.db,
		}
	}
}

type FavouriteRepositoryClient struct {
	q *gorm.DB
}

func (r *FavouriteRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *FavouriteRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *FavouriteRepositoryClient) CreateFavourite(ctx context.Context, data *domain.Favourite) error {
	if err := r.q.Debug().WithContext(ctx).Model(&domain.Favourite{}).Create(data).Error; err != nil {
		return errx.New(fiber.StatusInternalServerError, "failed to create favourite", err)
	}

	return nil
}

func (r *FavouriteRepositoryClient) DeleteFavourite(ctx context.Context, userID string, workerID string) error {
	result := r.q.Debug().WithContext(ctx).Model(&domain.Favourite{}).Where("worker_id = ? AND user_id = ?", workerID, userID).Delete(&domain.Favourite{})
	if result.Error != nil {
		return errx.New(fiber.StatusInternalServerError, "failed to delete favourite", result.Error)
	}
	if result.RowsAffected == 0 {
		return errx.New(fiber.StatusNotFound, "favourite not found", errors.New("favourite not found"))
	}

	return nil
}

func (r *FavouriteRepositoryClient) GetFavouriteByUserID(ctx context.Context, userID string) ([]domain.Favourite, error) {
	var favourite []domain.Favourite
	if err := r.q.Debug().WithContext(ctx).Model(&domain.Favourite{}).Where("user_id = ?", userID).Find(&favourite).Error; err != nil {
		return nil, errx.New(fiber.StatusInternalServerError, "failed to get favourite", err)
	}

	return favourite, nil
}
