package rest

import (
	"context"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/middleware"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type WorkerServiceHandler struct {
	service   port.WorkerServiceServiceItf
	token     port.TokenItf
	validator *validator.Validate
}

func NewWorkerServiceHandler(service port.WorkerServiceServiceItf, validator *validator.Validate, token port.TokenItf) *WorkerServiceHandler {
	return &WorkerServiceHandler{
		service:   service,
		validator: validator,
		token:     token,
	}
}

func (workerService *WorkerServiceHandler) Mount(router fiber.Router) {
	reviewRouter := router.Group("/worker-services")
	reviewRouter.Use(middleware.Request())
	reviewRouter.Use(middleware.Authentication(workerService.token, "worker"))

	reviewRouter.Post("", workerService.handleCreateWorkerService)
	reviewRouter.Get("", workerService.handleGetWorkerService)
}

func (workerService *WorkerServiceHandler) handleCreateWorkerService(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateWorkerServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := workerService.validator.Struct(req); err != nil {
		return err
	}

	if err := workerService.service.CreateWorkerService(ctx, req); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", ctx.Err())
	default:
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Worker Service Created",
		})
	}
}

func (workerService *WorkerServiceHandler) handleGetWorkerService(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	workerID, ok := c.Locals("userID").(string)
	if !ok {
		return errx.New(fiber.StatusUnauthorized, "invalid worker id from token", nil)
	}

	data, err := workerService.service.GetWorkerServicesByWorkerID(ctx, workerID)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", ctx.Err())
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": data,
		})
	}
}
