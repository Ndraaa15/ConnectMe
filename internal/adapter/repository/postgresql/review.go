package postgresql

import (
	"context"
	"errors"
	"net/http"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}

type ReviewRepository struct {
	db *gorm.DB
}

func (r *ReviewRepository) NewReviewRepositoryClient(tx bool) port.ReviewRepositoryClientItf {
	if tx {
		return &ReviewRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &ReviewRepositoryClient{
			q: r.db,
		}
	}
}

type ReviewRepositoryClient struct {
	q *gorm.DB
}

func (r *ReviewRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *ReviewRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *ReviewRepositoryClient) CreateReview(ctx context.Context, data *domain.Review) error {
	err := r.q.Debug().WithContext(ctx).Model(&domain.Review{}).Create(data).Error
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "foreign_key_violation":
				switch pqErr.Constraint {
				case "reviews_worker_id_fkey":
					return errx.New(http.StatusNotFound, "worker not found", err)
				case "reviews_user_id_fkey":
					return errx.New(http.StatusNotFound, "user not found", err)
				}
			case "unique_violation":
				if pqErr.Constraint == "reviews_pkey" {
					return errx.New(http.StatusConflict, "review already exists", err)
				}
			}
		}
		return errx.New(http.StatusInternalServerError, "failed to create review", err)
	}

	return nil
}
