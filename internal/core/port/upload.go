package port

import (
	"context"
	"mime/multipart"
)

type UploadServiceItf interface {
	Upload(ctx context.Context, file *multipart.FileHeader) (string, error)
}
