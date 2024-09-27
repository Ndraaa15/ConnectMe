package port

import (
	"context"
	"mime/multipart"
)

type StorageItf interface {
	UploadFile(ctx context.Context, file multipart.File) (string, error)
}
