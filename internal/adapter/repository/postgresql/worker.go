package postgresql

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"gorm.io/gorm"
)

func NewWorkerRepository(db *gorm.DB) *WorkerRepository {
	return &WorkerRepository{
		db: db,
	}
}

type WorkerRepository struct {
	db *gorm.DB
}

func (r *WorkerRepository) NewWorkerRepositoryClient(tx bool) port.WorkerRepositoryClientItf {
	if tx {
		return &WorkerRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &WorkerRepositoryClient{
			q: r.db,
		}
	}
}

type WorkerRepositoryClient struct {
	q *gorm.DB
}

func (r *WorkerRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *WorkerRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *WorkerRepositoryClient) GetWorkers(ctx context.Context) ([]domain.Worker, error) {
	var workers []domain.Worker

	if err := r.q.Debug().WithContext(ctx).Preload("WorkerServices").Preload("Tag").Find(&workers).Error; err != nil {
		return nil, err
	}

	return workers, nil
}
