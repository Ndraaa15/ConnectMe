package rest

import (
	"context"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/core/middleware"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type WorkerHandler struct {
	service   port.WorkerServiceItf
	token     port.TokenItf
	validator *validator.Validate
}

func NewWorkerHandler(service port.WorkerServiceItf, validator *validator.Validate, token port.TokenItf) *WorkerHandler {
	return &WorkerHandler{
		service:   service,
		validator: validator,
		token:     token,
	}
}

func (worker *WorkerHandler) Mount(router fiber.Router) {
	workerRouter := router.Group("/workers")
	workerRouter.Get(("/"), middleware.Request(), middleware.Authentication(worker.token), worker.GetWorkers)
	workerRouter.Get("/:id", middleware.Request(), middleware.Authentication(worker.token), worker.GetWorker)
}

func (worker *WorkerHandler) GetWorkers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var (
		err error
	)

	errChan := make(chan error, 1)
	resChan := make(chan interface{}, 1)

	go func() {
		res, err := worker.service.GetWorkers(ctx)
		if err != nil {
			errChan <- err
		}

		resChan <- res
	}()

	select {
	case err = <-errChan:
		return err
	case res := <-resChan:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "List workers",
			"workers": res,
		})
	}

}

func (worker *WorkerHandler) GetWorker(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var (
		err error
	)

	errChan := make(chan error, 1)
	resChan := make(chan interface{}, 1)

	go func() {
		workerID := c.Params("id")
		res, err := worker.service.GetWorker(ctx, workerID)
		if err != nil {
			errChan <- err
		}

		resChan <- res
	}()

	select {
	case err = <-errChan:
		return err
	case res := <-resChan:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Detail worker",
			"worker":  res,
		})
	}
}
