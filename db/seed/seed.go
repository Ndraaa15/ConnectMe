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

	RegisterSeeder()

	if name == "" {
		for seederName, seedFunc := range seeders {
			if err := seedFunc(db); err != nil {
				log.Fatal().Err(err).Msgf("Failed to seed %s", seederName)
			}
			log.Info().Msgf("Seed %s success", seederName)
		}
		return
	}

	seederFunc, exists := seeders[name]
	if !exists {
		log.Fatal().Msgf("Seeder %s not found", name)
		return
	}

	if err := seederFunc(db); err != nil {
		log.Fatal().Err(err).Msgf("Failed to seed %s", name)
		return
	}

}

func RegisterSeeder() {
	seeders["user"] = UserSeeder()
	seeders["tag"] = TagSeeder()
	seeders["worker"] = WorkerSeeder()
}
