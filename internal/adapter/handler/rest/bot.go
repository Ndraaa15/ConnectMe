package rest

import (
	"context"
	"errors"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/middleware"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BotHandler struct {
	service   port.BotServiceItf
	token     port.TokenItf
	validator *validator.Validate
}

func NewBotHandler(service port.BotServiceItf, token port.TokenItf, validator *validator.Validate) *BotHandler {
	return &BotHandler{
		service:   service,
		token:     token,
		validator: validator,
	}
}

func (bot *BotHandler) Mount(router fiber.Router) {
	botRouter := router.Group("/bots")
	botRouter.Use(middleware.Request())
	botRouter.Use(middleware.Authentication(bot.token, "user"))
	botRouter.Post("", bot.handleGenerateResponse)
}

func (bot *BotHandler) handleGenerateResponse(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userID, ok := c.Locals("userID").(string)
	if !ok {
		return fiber.ErrUnauthorized
	}

	problem := c.FormValue("problem")
	image, err := c.FormFile("image")
	if err != nil {
		if err.Error() != "there is no uploaded file associated with the given key" {
			return err
		}
	}

	content, err := bot.service.GenerateResponse(ctx, image, problem, userID)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  "Generate response success",
			"response": content,
		})
	}
}
