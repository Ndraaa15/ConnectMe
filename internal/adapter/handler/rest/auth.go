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

type AuthHandler struct {
	service   port.AuthServiceItf
	validator *validator.Validate
}

func NewAuthHandler(service port.AuthServiceItf, validator *validator.Validate) *AuthHandler {
	return &AuthHandler{
		service:   service,
		validator: validator,
	}
}

func (auth *AuthHandler) Mount(router fiber.Router) {
	authRouter := router.Group("/auth")
	authRouter.Post("/signup", middleware.Request(), auth.Register)
	authRouter.Post("/verify", middleware.Request(), auth.Verify)
	authRouter.Post("/signin", middleware.Request(), auth.Login)
}

func (auth *AuthHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var req dto.SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := auth.validator.Struct(req); err != nil {
		return err
	}

	res, err := auth.service.Register(ctx, req)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User created successfully",
			"id":      res,
		})
	}
}

func (auth *AuthHandler) Verify(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var req dto.VerifyAccountRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := auth.validator.Struct(req); err != nil {
		return err
	}

	res, err := auth.service.Verify(ctx, req)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User verified successfully",
			"id":      res,
		})
	}
}

func (auth *AuthHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var req dto.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := auth.validator.Struct(req); err != nil {
		return err
	}

	res, err := auth.service.Login(ctx, req)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User login successfully",
			"token":   res,
		})
	}
}
