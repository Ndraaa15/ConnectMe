package service

import (
	"context"
	"mime/multipart"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
)

type BotService struct {
	repository    port.BotRepositoryItf
	workerService port.WorkerServiceItf
	genai         port.GenaiItf
	storage       port.StorageItf
	cache         port.CacheItf
}

func NewBotService(repository port.BotRepositoryItf, genai port.GenaiItf, cache port.CacheItf, workerService port.WorkerServiceItf, storage port.StorageItf) *BotService {
	botService := &BotService{
		repository:    repository,
		genai:         genai,
		workerService: workerService,
		cache:         cache,
		storage:       storage,
	}

	return botService
}
func (bot *BotService) GenerateResponse(ctx context.Context, image *multipart.FileHeader, problem string, userID string) (dto.BotResponse, error) {
	repositoryClient := bot.repository.NewBotRepositoryClient(false)

	var (
		file      multipart.File
		fileBytes []byte
		photoUrl  string
		err       error
	)

	if image != nil {
		file, err = image.Open()
		if err != nil {
			return dto.BotResponse{}, err
		}
		defer file.Close()

		photoUrl, err = bot.storage.UploadFile(ctx, file)
		if err != nil {
			return dto.BotResponse{}, err
		}

	}

	content, err := bot.genai.GenerateResponseForProblem(ctx, problem, fileBytes)
	if err != nil {
		return dto.BotResponse{}, err
	}

	worker, err := bot.workerService.GetWorkersForBotResponse(ctx, content.Keyword)
	if err != nil {
		return dto.BotResponse{}, err
	}

	data := domain.Bot{
		UserID:   userID,
		Problem:  problem,
		Picture:  photoUrl,
		Solution: content.Solution,
	}

	err = repositoryClient.CreateBotResponse(ctx, &data)
	if err != nil {
		return dto.BotResponse{}, err
	}

	resp := dto.BotResponse{
		Problem:  problem,
		Image:    photoUrl,
		Solution: content.Solution,
		Worker:   worker,
	}

	return resp, nil
}
