package config

import (
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func NewFiber(conf env.App) *fiber.App {
	fiber := fiber.New(
		fiber.Config{
			AppName:          conf.Name,
			DisableKeepalive: false,
			Prefork:          false,
			StrictRouting:    false,
			ErrorHandler:     fiberErrorHandler(),
			RequestMethods:   []string{fiber.MethodGet, fiber.MethodHead, fiber.MethodPost, fiber.MethodPut, fiber.MethodDelete, fiber.MethodOptions, fiber.MethodPatch},
		},
	)

	return fiber
}

func fiberErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if ce, ok := err.(*errx.Errx); ok {
			log.Error().Err(ce).Msg(ce.Message)
			return c.Status(ce.Code).JSON(fiber.Map{
				"message": ce.Message,
				"error":   ce.Error(),
			})
		}

		if ve, ok := err.(validator.ValidationErrors); ok {
			out := make(map[string]string)
			for _, e := range ve {
				out[e.Field()] = util.GetErrorValidationMessage(e)
			}
			log.Error().Err(err).Msg("validation error")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Validation error",
				"error":   out,
			})
		}

		log.Error().Err(err).Msg("internal server error")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}
}
