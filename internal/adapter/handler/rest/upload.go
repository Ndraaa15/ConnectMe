package rest

import (
	"context"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/middleware"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UploadHandler struct {
	service   port.UploadServiceItf
	token     port.TokenItf
	validator *validator.Validate
}

func NewUploadServiceHandler(service port.UploadServiceItf, validator *validator.Validate, token port.TokenItf) *UploadHandler {
	return &UploadHandler{
		service:   service,
		validator: validator,
		token:     token,
	}
}

func (upload *UploadHandler) Mount(router fiber.Router) {
	uploadRouter := router.Group("/uploads")
	uploadRouter.Use(middleware.Request())
	uploadRouter.Use(middleware.Authentication(upload.token, "user", "worker"))

	uploadRouter.Post("", upload.handleUploadFile)
}

func (upload *UploadHandler) handleUploadFile(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	if err := upload.validator.Struct(file); err != nil {
		return err
	}

	res, err := upload.service.Upload(ctx, file)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", ctx.Err())
	default:
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Upload Success",
			"url":     res,
		})
	}
}
