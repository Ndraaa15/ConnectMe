package postgresql

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
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

func (r *OrderRepositoryClient) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {
	var order domain.Order
	err := r.q.Debug().WithContext(ctx).Where("id = ?", id).First(&order).Error
	return &order, err
}

func (r *OrderRepositoryClient) GetOrders(ctx context.Context) ([]*domain.Order, error) {
	var orders []*domain.Order
	err := r.q.Debug().WithContext(ctx).Find(&orders).Error
	return orders, err
}

func (r *OrderRepositoryClient) CreateAddressOrder(ctx context.Context, data *domain.AddressOrder) error {
	return r.q.Debug().WithContext(ctx).Model(&domain.AddressOrder{}).Create(data).Error
}
