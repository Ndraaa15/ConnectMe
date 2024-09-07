package service

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
)

type WorkerService struct {
	repository port.WorkerRepositoryItf
}

func NewWorkerService(repository port.WorkerRepositoryItf) *WorkerService {
	return &WorkerService{
		repository: repository,
	}
}

func (worker *WorkerService) GetWorkers(ctx context.Context) ([]domain.Worker, error) {
	repositoryClient := worker.repository.NewWorkerRepositoryClient(false)

	workers, err := repositoryClient.GetWorkers(ctx)
	if err != nil {
		return []domain.Worker{}, err
	}

	return workers, nil
}
