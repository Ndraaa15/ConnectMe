package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/app"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	app, err := app.NewApp()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create app")
	}

	app.RegisterHandler()
	app.Monitor()
	app.Health()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Run(); err != nil {
			log.Fatal().Err(err).Msg("Failed to run app")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigChan:
		log.Info().Msg("Received shutdown signal")
		cancel()
	case <-ctx.Done():
	}

	log.Info().Msg("Shutting down...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := app.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("Error during shutdown")
	}

	wg.Wait()
	log.Info().Msg("Shutdown complete")
}
