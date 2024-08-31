package middleware

import (
	"strings"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/paseto"
	"github.com/gofiber/fiber/v2"
)

func Authentication() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if authorization == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		authorizations := strings.SplitN(authorization, " ", 2)
		if len(authorizations) != 2 || authorizations[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		token := authorizations[1]
		paseto := paseto.NewPaseto()
		payload, err := paseto.Decode(token)
		if err != nil {
			return err
		}

		c.Locals("userID", payload.ID)

		return c.Next()
	}
}
