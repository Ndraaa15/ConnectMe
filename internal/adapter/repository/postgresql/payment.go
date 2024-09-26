package postgresql

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

type PaymentRepository struct {
	db *gorm.DB
}

func (r *PaymentRepository) NewPaymentRepositoryClient(tx bool) port.PaymentRepositoryClientItf {
	if tx {
		return &PaymentRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &PaymentRepositoryClient{
			q: r.db,
		}
	}
}

type PaymentRepositoryClient struct {
	q *gorm.DB
}

func (r *PaymentRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *PaymentRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *PaymentRepositoryClient) UpdatePayment(ctx context.Context, data *domain.Payment) error {
	if err := r.q.Debug().WithContext(ctx).Model(&domain.Payment{}).Where("id = ?", data.ID).Save(data).Error; err != nil {
		return errx.New(fiber.StatusInternalServerError, "failed to update payment", err)
	}

	return nil
}
