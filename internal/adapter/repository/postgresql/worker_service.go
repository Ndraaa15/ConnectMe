package postgresql

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
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

	if err != nil {
		return nil, errx.New(fiber.StatusInternalServerError, "failed to get worker services", err)
	}

	return workerServices, nil
}

func (r *WorkerServiceRepositoryClient) CreateWorkerService(ctx context.Context, data *domain.WorkerService) error {
	if err := r.q.Debug().WithContext(ctx).Model(&domain.WorkerService{}).Create(data).Error; err != nil {
		return errx.New(fiber.StatusInternalServerError, "failed to create worker service", err)
	}

	return nil
}
