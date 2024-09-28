package service

import (
	"context"
	"mime/multipart"

	"github.com/Ndraaa15/ConnectMe/internal/core/port"
)

type UploadService struct {
	storage port.StorageItf
}

func NewUploadService(storage port.StorageItf, cache port.CacheItf) *UploadService {
	return &UploadService{
		storage: storage,
	}
}

func (upload *UploadService) Upload(ctx context.Context, file *multipart.FileHeader) (string, error) {
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}

	defer fileContent.Close()

	fileUrl, err := upload.storage.UploadFile(ctx, fileContent)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
}
