package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
)

type WorkerServiceRepositoryItf interface {
	NewWorkerServiceRepositoryClient(tx bool) WorkerServiceRepositoryClientItf
}

type WorkerServiceRepositoryClientItf interface {
	Commit() error
	Rollback() error
	GetWorkerServicesByWorkerServiceIDs(ctx context.Context, workerServiceIDs []int64) ([]*domain.WorkerService, error)
	CreateWorkerService(ctx context.Context, data *domain.WorkerService) error
}

type WorkerServiceServiceItf interface {
	GetWorkerServicesByWorkerServiceIDs(ctx context.Context, workerServiceIDs []int64) ([]*domain.WorkerService, error)
	CreateWorkerService(ctx context.Context, req dto.CreateWorkerServiceRequest) error
}
