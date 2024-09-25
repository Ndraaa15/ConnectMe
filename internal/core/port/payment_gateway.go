package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
)

type PaymentGatewayItf interface {
	CreateTransaction(ctx context.Context, payment domain.Payment) (dto.TransactionResponse, error)
}
