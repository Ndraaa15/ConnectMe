package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
)

type FavouriteRepositoryItf interface {
	NewFavouriteRepositoryClient(tx bool) FavouriteRepositoryClientItf
}

type FavouriteRepositoryClientItf interface {
	Commit() error
	Rollback() error
	CreateFavourite(ctx context.Context, data *domain.Favourite) error
	DeleteFavourite(ctx context.Context, userID string, workerID string) error
	GetFavouriteByUserID(ctx context.Context, id string) ([]domain.Favourite, error)
}

type FavouriteServiceItf interface {
	CreateFavourite(ctx context.Context, req dto.CreateFavouriteRequest, userID string) error
	GetFavourites(ctx context.Context, userID string) ([]dto.WorkerResponse, error)
	DeleteFavourite(ctx context.Context, userID string, workerID string) error
}
