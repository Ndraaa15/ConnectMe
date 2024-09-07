package seed

import (
	"log"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func WorkerSeeder() Seeder {
	return func(db *gorm.DB) error {
		uuid1, err := uuid.Parse("079ca082-37c0-44d2-a648-f83bf482c4a3")
		if err != nil {
			log.Fatalf("Error parsing UUID: %v", err)
			return err
		}

		workers := []domain.Worker{
			{
				ID:             uuid1,
				TagID:          1,
				Description:    "Experienced plumber with over 5 years of work.",
				WorkExperience: 5,
				LowerPrice:     100.00,
				WorkerServices: []domain.WorkerService{
					{Service: "Leak Repair", Price: 150000, WorkerID: uuid1},
					{Service: "Pipe Installation", Price: 150000, WorkerID: uuid1},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/professional-plumber.jpg",
				WorkHour: pq.StringArray{"09:00", "17:00"},
			},
		}

		if err := db.CreateInBatches(&workers, len(workers)).Error; err != nil {
			return err
		}

		return nil
	}
}
