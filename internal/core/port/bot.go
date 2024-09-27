package port

import (
	"context"
	"mime/multipart"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
)

type BotRepositoryItf interface {
	NewBotRepositoryClient(tx bool) BotRepositoryClientItf
}

type BotRepositoryClientItf interface {
	Commit() error
	Rollback() error
	GetBotResponses(ctx context.Context) ([]domain.Bot, error)
	CreateBotResponse(ctx context.Context, data *domain.Bot) error
}

type BotServiceItf interface {
	GenerateResponse(ctx context.Context, image *multipart.FileHeader, problem string, userID string) (dto.BotResponse, error)
}
