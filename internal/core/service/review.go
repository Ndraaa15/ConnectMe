package service

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
)

type ReviewService struct {
	repository port.ReviewRepositoryItf
	cache      port.CacheItf
}

func NewReviewService(repository port.ReviewRepositoryItf, cache port.CacheItf) *ReviewService {
	reviewService := &ReviewService{
		repository: repository,
		cache:      cache,
	}

	return reviewService
}

func (s *ReviewService) CreateReview(ctx context.Context, req dto.CreateReviewRequest, userID string) error {
	repositoryClient := s.repository.NewReviewRepositoryClient(false)

	review := domain.Review{
		UserID:   userID,
		WorkerID: req.WorkerID,
		Rating:   req.Rating,
		Review:   req.Review,
	}

	err := repositoryClient.CreateReview(ctx, &review)
	if err != nil {
		return err
	}

	return nil
}
