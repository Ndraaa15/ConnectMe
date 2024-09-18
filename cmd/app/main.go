package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Ndraaa15/ConnectMe/db/migration"
	"github.com/Ndraaa15/ConnectMe/db/seed"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/app"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := env.NewEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create env")
		return
	}

	var wg sync.WaitGroup

	app, err := app.NewApp(env)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create app")
	}

	handleArgs(env)

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

func handleArgs(env *env.Env) {
	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	seedCmd := flag.NewFlagSet("seed", flag.ExitOnError)

	migrateAction := migrateCmd.String("action", "", "Specify 'up' or 'down' for migration")
	seedDomain := seedCmd.String("name", "", "Specify a domain for seeding (optional)")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			if err := migrateCmd.Parse(os.Args[2:]); err != nil {
				log.Fatal().Err(err).Msg("Failed to parse migrate command")
			}

			if *migrateAction == "" {
				log.Fatal().Msg("action is required")
			}
			migration.Migrate(env.Database, *migrateAction)
			os.Exit(1)
		case "seed":
			if err := seedCmd.Parse(os.Args[2:]); err != nil {
				log.Fatal().Err(err).Msg("Failed to parse seed command")
			}
			seed.Execute(env.Database, *seedDomain)
			os.Exit(1)
		}
	}
}
