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

type ReviewHandler struct {
	service   port.ReviewServiceItf
	token     port.TokenItf
	validator *validator.Validate
}

func NewReviewHandler(service port.ReviewServiceItf, validator *validator.Validate, token port.TokenItf) *ReviewHandler {
	return &ReviewHandler{
		service:   service,
		validator: validator,
		token:     token,
	}
}

func (review *ReviewHandler) Mount(router fiber.Router) {
	reviewRouter := router.Group("/reviews")
	reviewRouter.Use(middleware.Request())
	reviewRouter.Use(middleware.Authentication(review.token, "user"))

	reviewRouter.Post("", review.handleCreateReview)
}

func (review *ReviewHandler) handleCreateReview(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userID, ok := c.Locals("userID").(string)
	if !ok {
		return errx.New(fiber.StatusUnauthorized, "invalid user id from token", errors.New("UNAUTHORIZED"))
	}

	var req dto.CreateReviewRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := review.validator.Struct(req); err != nil {
		return err
	}

	if err := review.service.CreateReview(ctx, req, userID); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Review created successfully",
		})
	}
}
