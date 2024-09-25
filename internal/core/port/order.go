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
	CreateAddressOrder(ctx context.Context, data *domain.AddressOrder) error
}

type OrderServiceItf interface {
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest, userID string) (dto.PaymentResponse, error)
}
