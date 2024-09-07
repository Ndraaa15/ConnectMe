package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
)

type WorkerRepositoryItf interface {
	NewWorkerRepositoryClient(tx bool) WorkerRepositoryClientItf
}

type WorkerRepositoryClientItf interface {
	Commit() error
	Rollback() error
	GetWorkers(ctx context.Context) ([]domain.Worker, error)
}

type WorkerServiceItf interface {
	GetWorkers(ctx context.Context) ([]domain.Worker, error)
}
