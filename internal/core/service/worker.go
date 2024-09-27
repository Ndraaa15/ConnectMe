package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/util"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
)

type WorkerService struct {
	repository port.WorkerRepositoryItf
	cache      port.CacheItf
}

func NewWorkerService(repository port.WorkerRepositoryItf, cache port.CacheItf) *WorkerService {
	return &WorkerService{
		repository: repository,
		cache:      cache,
	}
}

func (worker *WorkerService) GetWorkers(ctx context.Context) ([]dto.WorkerResponse, error) {
	repositoryClient := worker.repository.NewWorkerRepositoryClient(false)

	data, err := repositoryClient.GetWorkers(ctx)
	if err != nil {
		return []dto.WorkerResponse{}, err
	}

	workerResponses := make([]dto.WorkerResponse, len(data))
	var wg sync.WaitGroup

	for i, worker := range data {
		wg.Add(1)
		go func(i int, worker domain.Worker) {
			defer wg.Done()
			formatWorkerResponse(&worker, &workerResponses[i])
			formatTagResponse(&worker.Tag, &workerResponses[i].Tag)
		}(i, worker)
	}

	wg.Wait()

	return workerResponses, nil
}

func (worker *WorkerService) GetWorkersByWorkerIDs(ctx context.Context, workerIDs []string) ([]dto.WorkerResponse, error) {
	fmt.Println("from worker service", workerIDs)
	repositoryClient := worker.repository.NewWorkerRepositoryClient(false)

	data, err := repositoryClient.GetWorkersByWorkerIDs(ctx, workerIDs)
	if err != nil {
		return []dto.WorkerResponse{}, err
	}

	workerResponses := make([]dto.WorkerResponse, len(data))
	var wg sync.WaitGroup

	for i, worker := range data {
		wg.Add(1)
		go func(i int, worker domain.Worker) {
			defer wg.Done()
			formatWorkerResponse(&worker, &workerResponses[i])
			formatTagResponse(&worker.Tag, &workerResponses[i].Tag)
		}(i, worker)
	}

	wg.Wait()

	return workerResponses, nil
}

func (worker *WorkerService) GetWorker(ctx context.Context, workerID string) (dto.WorkerDetailResponse, error) {
	// Todo : Adding field is available in time workhour
	repositoryClient := worker.repository.NewWorkerRepositoryClient(false)

	data, err := repositoryClient.GetWorker(ctx, workerID)
	if err != nil {
		return dto.WorkerDetailResponse{}, err
	}

	var workerResponse dto.WorkerDetailResponse
	formatWorkerDetailResponse(&data, &workerResponse)
	formatTagResponse(&data.Tag, &workerResponse.Tag)

	workerServices := make([]dto.WorkerServiceResponse, len(data.WorkerServices))
	workerReview := make([]dto.ReviewDetailResponse, len(data.Reviews))
	var wg sync.WaitGroup

	for i, workerService := range data.WorkerServices {
		wg.Add(1)
		go func(i int, workerService domain.WorkerService) {
			defer wg.Done()
			formatWorkerServiceResponse(&workerService, &workerServices[i])
		}(i, workerService)
	}

	wg.Wait()

	for i, review := range data.Reviews {
		wg.Add(1)
		go func(i int, review domain.Review) {
			defer wg.Done()
			formatReviewDetailResponse(&review, &workerReview[i])
		}(i, review)
	}

	wg.Wait()

	workerResponse.WorkerServices = workerServices
	workerResponse.Review = dto.ReviewResponse{
		Rating:        data.Rating,
		TotalRating:   data.TotalRating,
		ReviewsDetail: workerReview,
		TotalReview:   data.TotalReview,
	}

	return workerResponse, nil
}

func formatWorkerResponse(worker *domain.Worker, workerResp *dto.WorkerResponse) {
	*workerResp = dto.WorkerResponse{
		ID:         worker.ID,
		Name:       worker.Name,
		LowerPrice: worker.LowerPrice,
		Image:      worker.Image,
		Review:     dto.ReviewResponse{Rating: worker.Rating, TotalRating: worker.TotalRating},
	}
}

func formatWorkerDetailResponse(worker *domain.Worker, workerResp *dto.WorkerDetailResponse) {
	*workerResp = dto.WorkerDetailResponse{
		ID:             worker.ID,
		Name:           worker.Name,
		LowerPrice:     worker.LowerPrice,
		Image:          worker.Image,
		Description:    worker.Description,
		WorkExperience: worker.WorkExperience,
		WorkHour:       worker.WorkHour,
	}
}

func formatTagResponse(tag *domain.Tag, tagResp *dto.TagResponse) {
	*tagResp = dto.TagResponse{
		ID:             tag.ID,
		Tag:            tag.Tag,
		Specialization: tag.Specialization,
	}
}

func formatWorkerServiceResponse(workerService *domain.WorkerService, workerServiceResp *dto.WorkerServiceResponse) {
	*workerServiceResp = dto.WorkerServiceResponse{
		ID:      workerService.ID,
		Service: workerService.Service,
		Price:   workerService.Price,
	}
}

func formatReviewDetailResponse(review *domain.Review, reviewResp *dto.ReviewDetailResponse) {
	timeSent := util.GetTimeSinceCreation(review.CreatedAt)

	*reviewResp = dto.ReviewDetailResponse{
		Name:     review.User.FullName,
		Review:   review.Review,
		Rating:   review.Rating,
		TimeSent: timeSent,
	}
}
