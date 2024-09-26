package port

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
)

type ReviewRepositoryItf interface {
	NewReviewRepositoryClient(tx bool) ReviewRepositoryClientItf
}

type ReviewRepositoryClientItf interface {
	Commit() error
	Rollback() error
	CreateReview(ctx context.Context, data *domain.Review) error
}

type ReviewServiceItf interface {
	CreateReview(ctx context.Context, req dto.CreateReviewRequest, userID string) error
}
