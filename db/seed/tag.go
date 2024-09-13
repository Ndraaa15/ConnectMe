package seed

import (
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"gorm.io/gorm"
)

func TagSeeder() Seeder {
	return func(db *gorm.DB) error {
		tags := []domain.Tag{
			{
				ID:             1,
				Tag:            "Plumber",
				Specialization: "Plumbing",
			},
			{
				ID:             2,
				Tag:            "Electrician",
				Specialization: "Electrical",
			},
		}

		if err := db.Model(&domain.Tag{}).CreateInBatches(&tags, len(tags)).Error; err != nil {
			return err
		}

		return nil
	}
}
