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

type FavouriteHandler struct {
	service   port.FavouriteServiceItf
	validator *validator.Validate
	token     port.TokenItf
}

func NewFavouriteHandler(service port.FavouriteServiceItf, validator *validator.Validate, token port.TokenItf) *FavouriteHandler {
	return &FavouriteHandler{
		service:   service,
		validator: validator,
		token:     token,
	}
}

func (favourite *FavouriteHandler) Mount(router fiber.Router) {
	favouriteRouter := router.Group("/favourites")
	favouriteRouter.Use(middleware.Request())
	favouriteRouter.Use(middleware.Authentication(favourite.token, "user"))
	favouriteRouter.Post("", favourite.CreateFavourite)
	favouriteRouter.Get("", favourite.GetFavourites)
	favouriteRouter.Delete("/:id", favourite.DeleteFavourite)
}

func (h *FavouriteHandler) CreateFavourite(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userID, ok := c.Locals("userID").(string)
	if !ok {
		return errx.New(fiber.StatusUnauthorized, "invalid user id from token", errors.New("UNAUTHORIZED"))
	}

	var req dto.CreateFavouriteRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validator.Struct(req); err != nil {
		return err
	}

	if err := h.service.CreateFavourite(ctx, req, userID); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Favourite created successfully",
		})
	}
}

func (h *FavouriteHandler) GetFavourites(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userID, ok := c.Locals("userID").(string)
	if !ok {
		return errx.New(fiber.StatusUnauthorized, "invalid user id from token", errors.New("UNAUTHORIZED"))
	}

	res, err := h.service.GetFavourites(ctx, userID)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":    "List Favourites",
			"favourites": res,
		})
	}
}

func (h *FavouriteHandler) DeleteFavourite(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	userID, ok := c.Locals("userID").(string)
	if !ok {
		return errx.New(fiber.StatusUnauthorized, "invalid user id from token", errors.New("UNAUTHORIZED"))
	}

	workerID := c.Params("id")
	if err := h.service.DeleteFavourite(ctx, userID, workerID); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return errx.New(fiber.StatusRequestTimeout, "request timeout", errors.New("REQUEST TIMEOUT"))
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Favourite deleted successfully",
		})
	}
}
