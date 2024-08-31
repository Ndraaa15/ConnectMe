package rest

import (
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

func (auth *AuthHandler) Mount(srv *fiber.App) {
	a := srv.Group("/auth")
	a.Post("/login", auth.Register)
}

func (auth *AuthHandler) Register(c *fiber.Ctx) error {
	return nil
}

func (auth *AuthHandler) Login(c *fiber.Ctx) error {
	return nil
}
