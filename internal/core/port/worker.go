package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
)

type WorkerRepositoryItf interface {
	NewWorkerRepositoryClient(tx bool) WorkerRepositoryClientItf
}

type WorkerRepositoryClientItf interface {
	Commit() error
	Rollback() error
	GetWorkers(ctx context.Context) ([]domain.Worker, error)
	GetWorker(ctx context.Context, workerID string) (domain.Worker, error)
	GetWorkersByWorkerIDs(ctx context.Context, workerIDs []string) ([]domain.Worker, error)
}

type WorkerServiceItf interface {
	GetWorkers(ctx context.Context) ([]dto.WorkerResponse, error)
	GetWorker(ctx context.Context, workerID string) (dto.WorkerDetailResponse, error)
	GetWorkersByWorkerIDs(ctx context.Context, workerIDs []string) ([]dto.WorkerResponse, error)
}
