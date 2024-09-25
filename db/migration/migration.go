package migration

import (
	"github.com/Ndraaa15/ConnectMe/internal/adapter/config"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/rs/zerolog/log"
)

func Migrate(env env.Database, action string) {
	db := config.NewPostgreSQL(env)

	switch action {
	case "up":
		if err := db.AutoMigrate(
			&domain.User{},
			&domain.Worker{},
			&domain.WorkerService{},
			&domain.Tag{},
			&domain.Review{},
			&domain.Order{},
			&domain.AddressOrder{},
			&domain.Payment{},
		); err != nil {
			log.Fatal().Err(err).Msg("Failed to migrate")
		}
	case "down":
		if err := db.Migrator().DropTable(
			&domain.User{},
			&domain.Worker{},
			&domain.WorkerService{},
			&domain.Tag{},
			&domain.Review{},
			&domain.Order{},
			&domain.AddressOrder{},
			&domain.Payment{},
		); err != nil {
			log.Fatal().Err(err).Msg("Failed to drop")
		}
	}

	log.Info().Msgf("Migration %s success", action)
}
