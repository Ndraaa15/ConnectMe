package service

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
)

type WorkerServiceService struct {
	repository port.WorkerServiceRepositoryItf
	cache      port.CacheItf
}

func NewWorkerServiceService(repository port.WorkerServiceRepositoryItf, cache port.CacheItf) *WorkerServiceService {
	return &WorkerServiceService{
		repository: repository,
		cache:      cache,
	}
}

func (workerService *WorkerServiceService) GetWorkerServicesByWorkerServiceIDs(ctx context.Context, workerServiceIDs []int64) ([]*domain.WorkerService, error) {
	repositoryClient := workerService.repository.NewWorkerServiceRepositoryClient(false)

	data, err := repositoryClient.GetWorkerServicesByWorkerServiceIDs(ctx, workerServiceIDs)
	if err != nil {
		return nil, err
	}

	return data, nil
}
