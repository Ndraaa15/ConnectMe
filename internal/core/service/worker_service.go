package service

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
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

func (workerService *WorkerServiceService) CreateWorkerService(ctx context.Context, req dto.CreateWorkerServiceRequest) error {
	repositoryClient := workerService.repository.NewWorkerServiceRepositoryClient(false)

	data := domain.WorkerService{
		WorkerID: req.WorkerID,
		Service:  req.Service,
		Price:    req.Price,
	}

	err := repositoryClient.CreateWorkerService(ctx, &data)
	if err != nil {
		return err
	}

	return nil
}
