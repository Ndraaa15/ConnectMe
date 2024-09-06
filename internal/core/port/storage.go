package port

import "io"

type StorageItf interface {
	UploadFile(fileName string, fileBody io.Reader) error
	DeleteFile(fileName string) error
}
