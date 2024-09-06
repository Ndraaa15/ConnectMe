package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/config"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/db/migration"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/handler/rest"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/gomail"
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

func NewApp() (*App, error) {
	var app *App

	once.Do(func() {
		env, err := env.NewEnv()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create env")
			return
		}

		fiber := config.NewFiber(env.App)
		postgresql := config.NewPostgreSQL(env.Database)
		redis := config.NewRedis(env.Cache)
		validator := config.NewValidator()
		config.NewZerolog()

		err = migration.MigrateUp(postgresql)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to migrate")
			return
		}

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
	token := paseto.NewPaseto()
	email := gomail.NewGomail(a.env.Email)

	authRepository := postgresql.NewAuthRepository(a.db)
	authService := service.NewAuthService(authRepository, cache, token, email)
	authHandler := rest.NewAuthHandler(authService, a.validator)

	a.handlers = append(a.handlers, authHandler)
}

func (a *App) Run() error {
	log.Info().Msg("Starting Server...")
	v1 := a.srv.Group("/api/v1")
	for _, h := range a.handlers {
		h.Mount(v1)
	}

	a.srv.Use(middleware.Request())
	a.srv.Use(middleware.Cors())

	log.Info().Msg("Server is running")
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
		Title:   "MyService Metrics Page",
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
