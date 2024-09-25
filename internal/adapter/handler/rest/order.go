package rest

import (
	"context"
	"errors"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/middleware"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service   port.OrderServiceItf
	token     port.TokenItf
	validator *validator.Validate
}

func NewOrderHandler(service port.OrderServiceItf, validator *validator.Validate, token port.TokenItf) *OrderHandler {
	return &OrderHandler{
		service:   service,
		validator: validator,
		token:     token,
	}
}

func (order *OrderHandler) Mount(router fiber.Router) {
	orderRouter := router.Group("/orders")
	orderRouter.Use(middleware.Request())
	orderRouter.Use(middleware.Authentication(order.token, "user"))
	orderRouter.Post("", middleware.Request(), order.CreateOrder)
}

func (order *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userID, ok := c.Locals("userID").(string)
	if !ok {
		return errx.New(fiber.StatusUnauthorized, "invalid user id from token", errors.New("UNAUTHORIZED"))
	}

	var (
		err error
	)

	errChan := make(chan error, 1)
	resChan := make(chan interface{}, 1)

	go func() {
		var req dto.CreateOrderRequest
		if err := c.BodyParser(&req); err != nil {
			errChan <- err
		}

		if err := order.validator.Struct(req); err != nil {
			errChan <- err
		}

		res, err := order.service.CreateOrder(ctx, req, userID)
		if err != nil {
			errChan <- err
		}

		resChan <- res
	}()

	select {
	case err = <-errChan:
		return err
	case res := <-resChan:
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":     "Order Created",
			"transaction": res,
		})
	}
}
