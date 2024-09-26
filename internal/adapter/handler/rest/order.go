package rest

import (
	"context"
	"errors"
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
	orderRouter.Patch("/:id", order.UpdateOrder)
}

func (order *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userID, ok := c.Locals("userID").(string)
	if !ok {
		return errx.New(fiber.StatusUnauthorized, "invalid user id from token", errors.New("UNAUTHORIZED"))
	}

	var req dto.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := order.validator.Struct(req); err != nil {
		return err
	}

	res, err := order.service.CreateOrder(ctx, req, userID)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
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

	filter, err := parseFilterGetOrders(c)
	if err != nil {
		return err
	}

	orders, err := r.service.GetOrders(ctx, userID, filter)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Orders Fetched",
			"orders":  orders,
		})
	}
}

func (r *OrderHandler) GetOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var (
		err error
	)

	orderID := c.Params("id")

	order, err := r.service.GetOrder(ctx, orderID)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Order Fetched",
			"order":   order,
		})
	}
}

func (r *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	orderID := c.Params("id")

	var req dto.UpdateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := r.validator.Struct(req); err != nil {
		return err
	}

	if err := r.service.UpdateOrder(ctx, orderID, req); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Order Updated",
		})
	}
}

func parseFilterGetOrders(c *fiber.Ctx) (dto.GetOrderFilter, error) {
	filter := dto.GetOrderFilter{}

	if statusQuery := c.Query("status"); statusQuery != "" {
		statusesQuery := strings.Split(statusQuery, ",")
		statuses := []domain.StatusOrder{}
		for _, s := range statusesQuery {
			s, err := domain.ParseStatusOrder(s)
			if err != nil {
				return dto.GetOrderFilter{}, err
			}
			statuses = append(statuses, s)
		}
		filter.Status = statuses
	}

	return filter, nil
}
