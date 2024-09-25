package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
)

type PaymentRepositoryItf interface {
	NewPaymentRepositoryClient(tx bool) PaymentRepositoryClientItf
}

type PaymentRepositoryClientItf interface {
	Commit() error
	Rollback() error
	CreatePayment(ctx context.Context, data *domain.Payment) error
}

type PaymentServiceItf interface {
}
