package seed

import (
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ReviewSeeder() Seeder {
	return func(db *gorm.DB) error {
		reviews := []domain.Review{
			{
				ID:          1,
				WorkerID:    uuid.MustParse("079ca082-37c0-44d2-a648-f83bf482c4a3"),
				UserID:      uuid.MustParse("3eae758c-7c50-4266-b28d-eb180ffd8336"),
				Rating:      4.5,
				Description: "He is very professional and fast in fixing the leak in my house. I will definitely call him again if I have another problem.",
			},
			{
				ID:          2,
				WorkerID:    uuid.MustParse("fa35118b-1c30-48cb-a2de-2ccb3fa3281c"),
				UserID:      uuid.MustParse("3eae758c-7c50-4266-b28d-eb180ffd8336"),
				Rating:      4.3,
				Description: "He is very professional and fast in fixing the wiring in my house. I will definitely call him again if I have another problem.",
			},
			{
				ID:          3,
				WorkerID:    uuid.MustParse("079ca082-37c0-44d2-a648-f83bf482c4a3"),
				UserID:      uuid.MustParse("51487321-5387-4204-a788-686f7b1b80d4"),
				Rating:      3.5,
				Description: "He is very professional and fast in fixing the leak in my house. I will definitely call him again if I have another problem.",
			},
			{
				ID:          4,
				WorkerID:    uuid.MustParse("fa35118b-1c30-48cb-a2de-2ccb3fa3281c"),
				UserID:      uuid.MustParse("51487321-5387-4204-a788-686f7b1b80d4"),
				Rating:      4.0,
				Description: "He is very professional and fast in fixing the wiring in my house. I will definitely call him again if I have another problem.",
			},
		}

		if err := db.Model(&domain.Review{}).CreateInBatches(&reviews, len(reviews)).Error; err != nil {
			return err
		}

		return nil
	}
}
