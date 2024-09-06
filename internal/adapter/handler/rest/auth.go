package rest

import (
	"context"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
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
	authRouter.Post("/signup", auth.Register)
}

func (auth *AuthHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	var (
		err error
	)

	errChan := make(chan error, 1)
	resChan := make(chan interface{}, 1)

	go func() {
		var req dto.SignUpRequest
		if err := c.BodyParser(&req); err != nil {
			errChan <- err
		}

		if err := auth.validator.Struct(req); err != nil {
			errChan <- err
		}

		res, err := auth.service.Register(ctx, req)
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
			"message": "User created successfully",
			"id":      res,
		})
	}
}

func (auth *AuthHandler) Verify(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	var (
		err error
	)

	errChan := make(chan error, 1)
	resChan := make(chan interface{}, 1)

	go func() {
		var req dto.VerifyAccountRequest
		if err := c.BodyParser(&req); err != nil {
			errChan <- err
		}

		if err := auth.validator.Struct(req); err != nil {
			errChan <- err
		}

		res, err := auth.service.Verify(ctx, req)
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
			"message": "User verified successfully",
			"id":      res,
		})
	}
}
