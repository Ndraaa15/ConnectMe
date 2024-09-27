package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Authentication(tokenSvc port.TokenItf, role ...string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if authorization == "" {
			return errx.New(fiber.StatusUnauthorized, "missing token", errors.New("UNAUTHORIZED"))
		}

		authorizations := strings.SplitN(authorization, " ", 2)
		if len(authorizations) != 2 || authorizations[0] != "Bearer" {
			return errx.New(fiber.StatusUnauthorized, "invalid token", errors.New("UNAUTHORIZED"))
		}

		token := authorizations[1]
		payload, err := tokenSvc.Decode(token)
		if err != nil {
			log.Error().Err(err).Msg("Failed to decode token")
			return err
		}
		fmt.Println("payload", payload)
		if len(role) > 0 {
			isRole := false
			for _, r := range role {
				if r == payload.Role {
					isRole = true
					break
				}
			}

			if !isRole {
				return errx.New(fiber.StatusUnauthorized, "unauthorized", errors.New("UNAUTHORIZED"))
			}
		}

		c.Locals("userID", payload.ID)

		return c.Next()
	}
}
