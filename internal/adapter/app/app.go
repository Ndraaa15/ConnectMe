package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/config"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/handler/rest"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/cloudinary"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/gemini"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/gomail"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/midtrans"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/paseto"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/repository/cache"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/repository/postgresql"
	"github.com/Ndraaa15/ConnectMe/internal/core/middleware"
	"github.com/Ndraaa15/ConnectMe/internal/core/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var once sync.Once

type App struct {
	env       *env.Env
	srv       *fiber.App
	db        *gorm.DB
	cache     *redis.Client
	validator *validator.Validate
	handlers  []Handler
}

type Handler interface {
	Mount(srv fiber.Router)
}

func NewApp(env *env.Env) (*App, error) {
	var app *App

	once.Do(func() {
		fiber := config.NewFiber(env.App)
		postgresql := config.NewPostgreSQL(env.Database)
		redis := config.NewRedis(env.Cache)
		validator := config.NewValidator()
		config.NewZerolog()

		app = &App{
			env:       env,
			srv:       fiber,
			db:        postgresql,
			validator: validator,
			cache:     redis,
		}
	})

	return app, nil
}

func (a *App) RegisterHandler() {
	cache := cache.NewRedisClient(a.cache)
	token := paseto.NewPaseto(a.env.Token)
	email := gomail.NewGomail(a.env.Email)
	paymentGateway := midtrans.NewMidtrans(a.env.PaymentGateway)
	genai := gemini.NewGemini(a.env.Gemini)
	storage := cloudinary.NewCloudinary(a.env.Storage)

	authRepository := postgresql.NewAuthRepository(a.db)
	authService := service.NewAuthService(authRepository, cache, token, email)
	authHandler := rest.NewAuthHandler(authService, a.validator)

	workerRepository := postgresql.NewWorkerRepository(a.db)
	workerService := service.NewWorkerService(workerRepository, cache)
	workerHandler := rest.NewWorkerHandler(workerService, a.validator, token)

	workerServiceRepository := postgresql.NewWorkerServiceRepository(a.db)
	workerServiceService := service.NewWorkerServiceService(workerServiceRepository, cache)

	order := postgresql.NewOrderRepository(a.db)
	orderService := service.NewOrderService(order, cache, workerServiceService, paymentGateway)
	orderHandler := rest.NewOrderHandler(orderService, a.validator, token)

	reviewRepository := postgresql.NewReviewRepository(a.db)
	reviewService := service.NewReviewService(reviewRepository, cache)
	reviewHandler := rest.NewReviewHandler(reviewService, a.validator, token)

	favouriteRepository := postgresql.NewFavouriteRepository(a.db)
	favouriteService := service.NewFavouriteService(favouriteRepository, workerService, cache)
	favouriteHandler := rest.NewFavouriteHandler(favouriteService, a.validator, token)

	botRepository := postgresql.NewBotRepository(a.db)
	botService := service.NewBotService(botRepository, genai, cache, workerService, storage)
	botHandler := rest.NewBotHandler(botService, token, a.validator)

	a.handlers = append(a.handlers, authHandler, workerHandler, orderHandler, reviewHandler, favouriteHandler, botHandler)
}

func (a *App) Run() error {
	log.Info().Msg("Starting Server...")
	v1 := a.srv.Group("/api/v1")
	for _, h := range a.handlers {
		h.Mount(v1)
	}

	a.srv.Use(middleware.Request())
	a.srv.Use(middleware.Cors())

	log.Info().Msgf("Server is running on %s:%s", a.env.App.Address, a.env.App.Port)
	return a.srv.Listen(fmt.Sprintf("%s:%s", a.env.App.Address, a.env.App.Port))
}

func (a *App) Shutdown(ctx context.Context) error {
	done := make(chan struct{})

	go func() {
		defer close(done)

		if err := a.srv.Shutdown(); err != nil {
			log.Err(err).Msg("Error While Server Shutdown")
		}

		db, err := a.db.DB()
		if err != nil {
			log.Err(err).Msg("Error While Initialize SQL")
		}

		if err := db.Close(); err != nil {
			log.Err(err).Msg("Error While Database Shutdown")
		}

		if err := a.cache.Close(); err != nil {
			log.Err(err).Msg("Error While Cache Shutdown")
		}

		time.Sleep(2 * time.Second)
	}()

	select {
	case <-done:
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

func (a *App) Monitor() {
	a.srv.Get("/metrics", monitor.New(monitor.Config{
		Title:   "ConnectMe Metrics",
		Refresh: 5 * time.Second,
	}))
}

func (a *App) Health() {
	a.srv.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "health 💅",
		})
	})
}
