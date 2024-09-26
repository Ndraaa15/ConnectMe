package postgresql

import (
	"context"
	"errors"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

type OrderRepository struct {
	db *gorm.DB
}

func (r *OrderRepository) NewOrderRepositoryClient(tx bool) port.OrderRepositoryClientItf {
	if tx {
		return &OrderRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &OrderRepositoryClient{
			q: r.db,
		}
	}
}

type OrderRepositoryClient struct {
	q *gorm.DB
}

func (r *OrderRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *OrderRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *OrderRepositoryClient) CreateOrder(ctx context.Context, data *domain.Order) error {
	return r.q.Debug().WithContext(ctx).Model(&domain.Order{}).Create(data).Error
}

func (r *OrderRepositoryClient) GetOrderByID(ctx context.Context, id string) (domain.Order, error) {
	var order domain.Order
	err := r.q.Debug().WithContext(ctx).Preload("Worker.Tag").Preload("Worker.WorkerServices").Preload("Payment").Preload("Address").Model(&domain.Order{}).Where("order_id = ?", id).First(&order).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Order{}, errx.New(fiber.StatusNotFound, "order not found", err)
		}
		return domain.Order{}, errx.New(fiber.StatusInternalServerError, "failed to get order", err)
	}

	return order, nil
}

func (r *OrderRepositoryClient) GetOrdersByUserID(ctx context.Context, userID string, filter dto.GetOrderFilter) ([]domain.Order, error) {
	var orders []domain.Order

	queryBuilder := r.q.Debug().WithContext(ctx).Preload("Worker.Tag").Preload("Payment").Model(&domain.Order{}).Where("user_id = ?", userID)

	if filter.Status != nil {
		queryBuilder = queryBuilder.Where("order_status IN ?", filter.Status)
	}

	err := queryBuilder.Find(&orders).Error
	if err != nil {
		return []domain.Order{}, errx.New(fiber.StatusInternalServerError, "failed to get orders", err)
	}

	return orders, nil
}

func (r *OrderRepositoryClient) UpdateOrder(ctx context.Context, data *domain.Order) error {
	if err := r.q.Debug().WithContext(ctx).Model(&domain.Order{}).Where("order_id = ?", data.OrderID).Updates(data).Error; err != nil {
		return errx.New(fiber.StatusInternalServerError, "failed to update order", err)
	}

	return nil
}
