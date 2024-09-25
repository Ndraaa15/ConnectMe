package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
)

type OrderRepositoryItf interface {
	NewOrderRepositoryClient(tx bool) OrderRepositoryClientItf
}

type OrderRepositoryClientItf interface {
	Commit() error
	Rollback() error
	CreateOrder(ctx context.Context, data *domain.Order) error
	GetOrdersByUserID(ctx context.Context, userID string, filter dto.GetOrderFilter) ([]domain.Order, error)
	GetOrderByID(ctx context.Context, id string) (domain.Order, error)
	UpdateOrder(ctx context.Context, data *domain.Order) error
}

type OrderServiceItf interface {
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest, userID string) (dto.TransactionResponse, error)
	GetOrders(ctx context.Context, userID string, filter dto.GetOrderFilter) ([]dto.OrderResponse, error)
	GetOrder(ctx context.Context, orderID string) (dto.OrderDetailResponse, error)
	UpdateOrder(ctx context.Context, orderID string, req dto.UpdateOrderRequest) error
}
