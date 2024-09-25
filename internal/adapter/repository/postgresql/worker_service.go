package postgresql

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"gorm.io/gorm"
)

func NewWorkerServiceRepository(db *gorm.DB) *WorkerServiceRepository {
	return &WorkerServiceRepository{
		db: db,
	}
}

type WorkerServiceRepository struct {
	db *gorm.DB
}

func (r *WorkerServiceRepository) NewWorkerServiceRepositoryClient(tx bool) port.WorkerServiceRepositoryClientItf {
	if tx {
		return &WorkerServiceRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &WorkerServiceRepositoryClient{
			q: r.db,
		}
	}
}

type WorkerServiceRepositoryClient struct {
	q *gorm.DB
}

func (r *WorkerServiceRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *WorkerServiceRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *WorkerServiceRepositoryClient) GetWorkerServicesByWorkerServiceIDs(ctx context.Context, workerServiceIDs []int64) ([]*domain.WorkerService, error) {
	var workerServices []*domain.WorkerService
	err := r.q.Debug().WithContext(ctx).Where("id IN ?", workerServiceIDs).Find(&workerServices).Error
	return workerServices, err
}
