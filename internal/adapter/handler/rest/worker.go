package rest

import (
	"context"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type WorkerHandler struct {
	service   port.WorkerServiceItf
	validator *validator.Validate
}

func NewWorkerHandler(service port.WorkerServiceItf, validator *validator.Validate) *WorkerHandler {
	return &WorkerHandler{
		service:   service,
		validator: validator,
	}
}

func (worker *WorkerHandler) Mount(router fiber.Router) {
	workerRouter := router.Group("/workers")
	workerRouter.Get("/", worker.GetWorkers)
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
