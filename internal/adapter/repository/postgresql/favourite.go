package postgresql

import (
	"context"

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
		return &ReviewRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &ReviewRepositoryClient{
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

func (r *FavouriteRepositoryClient) DeleteFavourite(ctx context.Context, data *domain.Favourite) error {
	if err := r.q.Debug().WithContext(ctx).Model(&domain.Favourite{}).Where("worker_id = ? AND user_id = ?", data.WorkerID, data.UserID).Delete(data).Error; err != nil {
		return errx.New(fiber.StatusInternalServerError, "failed to delete favourite", err)
	}

	return nil
}

func (r *FavouriteRepositoryClient) GetFavouriteByUserID(ctx context.Context, id string) ([]domain.Favourite, error) {
	var favourite []domain.Favourite
	if err := r.q.Debug().WithContext(ctx).Model(&domain.Favourite{}).Where("user_id = ?", id).Find(&favourite).Error; err != nil {
		return nil, errx.New(fiber.StatusInternalServerError, "failed to get favourite", err)
	}

	return favourite, nil
}
