package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Request() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		log.Info().Msgf("Request: %s %s %s", c.IP(), c.Method(), c.Path())
		return c.Next()
	}
}
