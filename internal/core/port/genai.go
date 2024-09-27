package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
)

type GenaiItf interface {
	GenerateResponseForProblem(ctx context.Context, text string, picture []byte) (dto.ResponseProblem, error)
}
