package seed

import (
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func WorkerSeeder() Seeder {
	return func(db *gorm.DB) error {
		workers := []domain.Worker{
			{
				ID:             "079ca082-37c0-44d2-a648-f83bf482c4a3",
				Name:           "John Doe",
				TagID:          1,
				Description:    "Experienced plumber with over 5 years of work.",
				WorkExperience: 5,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Leak Repair",
						Price:    150000,
						WorkerID: uuid.MustParse("079ca082-37c0-44d2-a648-f83bf482c4a3"),
					},
					{
						Service:  "Pipe Installation",
						Price:    160000,
						WorkerID: uuid.MustParse("079ca082-37c0-44d2-a648-f83bf482c4a3"),
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/professional-plumber.jpg",
				WorkHour: pq.StringArray{"09:00", "17:00", "18:00", "19:00", "20:00"},
			},
			{
				ID:             "fa35118b-1c30-48cb-a2de-2ccb3fa3281c",
				Name:           "Jack Doe",
				TagID:          2,
				Description:    "Experienced electrician with over 5 years of work.",
				WorkExperience: 5,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Wiring Installation",
						Price:    100000,
						WorkerID: uuid.MustParse("fa35118b-1c30-48cb-a2de-2ccb3fa3281c"),
					},
					{
						Service:  "Electrical Repair",
						Price:    150000,
						WorkerID: uuid.MustParse("fa35118b-1c30-48cb-a2de-2ccb3fa3281c"),
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/What-Can-I-Do-as-an-Electrician.jpg",
				WorkHour: pq.StringArray{"09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00"},
			},
		}

		if err := db.Model(&domain.Worker{}).CreateInBatches(&workers, len(workers)).Error; err != nil {
			return err
		}

		return nil
	}
}
