package service

import "github.com/Ndraaa15/ConnectMe/internal/core/port"

type FavouriteService struct {
	repository port.ReviewRepositoryItf
	cache      port.CacheItf
}

func NewFavouriteService(repository port.ReviewRepositoryItf, cache port.CacheItf) *ReviewService {
	reviewService := &ReviewService{
		repository: repository,
		cache:      cache,
	}

	return reviewService
}
