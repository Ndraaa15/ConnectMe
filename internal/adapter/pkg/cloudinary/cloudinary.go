package cloudinary

import (
	"context"
	"mime/multipart"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/rs/zerolog/log"
)

type Cloudinary struct {
	cloudinary *cloudinary.Cloudinary
	folder     string
}

func NewCloudinary(conf env.Storage) *Cloudinary {
	cld, err := cloudinary.NewFromParams(conf.Name, conf.ApiKey, conf.ApiSecret)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize cloudinary")
	}
	return &Cloudinary{
		cloudinary: cld,
		folder:     conf.Folder,
	}
}

func (c *Cloudinary) UploadFile(ctx context.Context, file multipart.File) (string, error) {
	unique := true
	res, err := c.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:         c.folder,
		UniqueFilename: unique,
	})

	if err != nil {
		return "", err
	}

	return res.SecureURL, nil
}
