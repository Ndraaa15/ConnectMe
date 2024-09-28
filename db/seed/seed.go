package seed

import (
	"github.com/Ndraaa15/ConnectMe/internal/adapter/config"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Seeder func(db *gorm.DB) error

var seeders = map[string]Seeder{}

func Execute(env env.Database, name string) {
	db := config.NewPostgreSQL(env)

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatal().Msgf("Transaction failed and rolled back due to panic: %v", r)
		}
	}()

	RegisterSeeder()

	if name == "" {
		for seederName, seedFunc := range seeders {
			if err := seedFunc(tx); err != nil {
				tx.Rollback()
				log.Fatal().Err(err).Msgf("Failed to seed %s, transaction rolled back", seederName)
				return
			}
			log.Info().Msgf("Seed %s success", seederName)
		}

		if err := tx.Commit().Error; err != nil {
			log.Fatal().Err(err).Msg("Failed to commit transaction")
		}
		return
	}

	seederFunc, exists := seeders[name]
	if !exists {
		log.Fatal().Msgf("Seeder %s not found", name)
		return
	}

	if err := seederFunc(tx); err != nil {
		tx.Rollback()
		log.Fatal().Err(err).Msgf("Failed to seed %s, transaction rolled back", name)
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}
}

func RegisterSeeder() {
	seeders["user"] = UserSeeder()
	seeders["tag"] = TagSeeder()
	seeders["worker"] = WorkerSeeder()
	seeders["review"] = ReviewSeeder()
	seeders["order"] = OrderSeeder()
	seeders["payment"] = PaymentSeeder()
}
