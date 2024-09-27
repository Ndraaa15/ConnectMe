package postgresql

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewBotRepository(db *gorm.DB) *BotRepository {
	return &BotRepository{
		db: db,
	}
}

type BotRepository struct {
	db *gorm.DB
}

func (r *BotRepository) NewBotRepositoryClient(tx bool) port.BotRepositoryClientItf {
	if tx {
		return &BotRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &BotRepositoryClient{
			q: r.db,
		}
	}
}

type BotRepositoryClient struct {
	q *gorm.DB
}

func (r *BotRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *BotRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}

func (r *BotRepositoryClient) CreateBotResponse(ctx context.Context, data *domain.Bot) error {
	if err := r.q.Debug().WithContext(ctx).Model(&domain.Bot{}).Create(data).Error; err != nil {
		return errx.New(fiber.StatusInternalServerError, "failed create bot response", err)
	}

	return nil
}

func (r *BotRepositoryClient) GetBotResponses(ctx context.Context) ([]domain.Bot, error) {
	var reponsesBot []domain.Bot

	if err := r.q.Debug().WithContext(ctx).Find(&reponsesBot).Error; err != nil {
		return nil, errx.New(fiber.StatusInternalServerError, "failed to get response bot", err)
	}

	return reponsesBot, nil
}
