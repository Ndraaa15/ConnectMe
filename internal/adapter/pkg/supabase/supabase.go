package supabase

import (
	"fmt"
	"io"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	storage_go "github.com/supabase-community/storage-go"
)

type Supabase struct {
	SupabaseClient *storage_go.Client
	Bucket         string
}

func NewSupabase(conf *env.Storage) *Supabase {
	url := fmt.Sprintf("https://%s.supabase.co/storage/v1", conf.ProjectID)
	client := storage_go.NewClient(url, conf.ApiKey, nil)

	return &Supabase{
		SupabaseClient: client,
		Bucket:         conf.Bucket,
	}
}

func (s *Supabase) UploadFile(fileName string, fileBody io.Reader) error {
	_, err := s.SupabaseClient.UploadFile(s.Bucket, fileName, fileBody)
	if err != nil {
		return err
	}

	return nil
}

func (s *Supabase) DeleteFile(fileName string) error {
	_, err := s.SupabaseClient.RemoveFile(s.Bucket, []string{fileName})
	if err != nil {
		return err
	}

	return nil
}
