package rest

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
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
	orderRouter.Post("", order.CreateOrder)
	orderRouter.Get("", order.GetOrders)
	orderRouter.Get("/:id", order.GetOrder)
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
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	case err = <-errChan:
		return err
	case res := <-resChan:
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":     "Order Created",
			"transaction": res,
		})
	}
}

func (r *OrderHandler) GetOrders(c *fiber.Ctx) error {
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
		filter, err := parseFilterGetOrders(c)
		if err != nil {
			errChan <- err
		}

		fmt.Println("filter", filter)

		orders, err := r.service.GetOrders(ctx, userID, filter)
		if err != nil {
			errChan <- err
		}

		resChan <- orders
	}()

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	case err = <-errChan:
		return err
	case res := <-resChan:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Orders Fetched",
			"orders":  res,
		})
	}
}

func (r *OrderHandler) GetOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var (
		err error
	)

	errChan := make(chan error, 1)
	resChan := make(chan interface{}, 1)

	orderID := c.Params("id")
	go func() {
		order, err := r.service.GetOrder(ctx, orderID)
		if err != nil {
			errChan <- err
		}

		resChan <- order
	}()

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	case err = <-errChan:
		return err
	case res := <-resChan:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Order Fetched",
			"order":   res,
		})
	}
}

func parseFilterGetOrders(c *fiber.Ctx) (dto.GetOrderFilter, error) {
	filter := dto.GetOrderFilter{}

	if statusQuery := c.Query("status"); statusQuery != "" {
		statusesQuery := strings.Split(statusQuery, ",")
		statuses := []domain.StatusOrder{}
		for _, s := range statusesQuery {
			s, err := parseStatusOrder(s)
			if err != nil {
				return dto.GetOrderFilter{}, err
			}
			statuses = append(statuses, s)
		}
		filter.Status = statuses
	}

	return filter, nil
}

func parseStatusOrder(status string) (domain.StatusOrder, error) {
	switch status {
	case "on_going":
		return domain.StatusOrderOnGoing, nil
	case "completed":
		return domain.StatusOrderFinished, nil
	case "canceled":
		return domain.StatusOrderCanceled, nil
	default:
		return domain.StatusOrderUnknown, errx.New(fiber.StatusBadRequest, "invalid status order", errors.New("invalid status order"))
	}
}
