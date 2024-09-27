package service

import (
	"context"
	"fmt"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
)

type FavouriteService struct {
	repository    port.FavouriteRepositoryItf
	workerService port.WorkerServiceItf
	cache         port.CacheItf
}

func NewFavouriteService(repository port.FavouriteRepositoryItf, workerService port.WorkerServiceItf, cache port.CacheItf) *FavouriteService {
	reviewService := &FavouriteService{
		repository:    repository,
		workerService: workerService,
		cache:         cache,
	}

	return reviewService
}

func (favourite *FavouriteService) CreateFavourite(ctx context.Context, req dto.CreateFavouriteRequest, userID string) error {
	repositoryClient := favourite.repository.NewFavouriteRepositoryClient(false)

	data := domain.Favourite{
		UserID:   userID,
		WorkerID: req.WorkerID,
	}

	if err := repositoryClient.CreateFavourite(ctx, &data); err != nil {
		return err
	}

	return nil
}

func (favourite *FavouriteService) GetFavourites(ctx context.Context, userID string) ([]dto.WorkerResponse, error) {
	repositoryClient := favourite.repository.NewFavouriteRepositoryClient(false)

	favourites, err := repositoryClient.GetFavouriteByUserID(ctx, userID)
	if err != nil {
		return []dto.WorkerResponse{}, err
	}

	workerIDs := make([]string, 0)
	for _, favourite := range favourites {
		workerIDs = append(workerIDs, favourite.WorkerID)
	}

	data, err := favourite.workerService.GetWorkersByWorkerIDs(ctx, workerIDs)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (favourite *FavouriteService) DeleteFavourite(ctx context.Context, userID string, workerID string) error {
	repositoryClient := favourite.repository.NewFavouriteRepositoryClient(false)

	if err := repositoryClient.DeleteFavourite(ctx, userID, workerID); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
