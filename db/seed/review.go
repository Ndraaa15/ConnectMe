package seed

import (
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"gorm.io/gorm"
)

func ReviewSeeder() Seeder {
	return func(db *gorm.DB) error {
		reviews := []domain.Review{
			{
				WorkerID: "69f74995-fd37-4f32-8583-e8baad2cd18f",
				UserID:   "3eae758c-7c50-4266-b28d-eb180ffd8336",
				Rating:   4.5,
				Review:   "Sangat profesional dan cepat dalam memperbaiki kabel di rumah saya. Pasti akan memanggilnya lagi jika ada masalah lain.",
			},
			{
				WorkerID: "69f74995-fd37-4f32-8583-e8baad2cd18f",
				UserID:   "d296f037-e55d-400f-8fb6-21dc97cc8fad",
				Rating:   3.8,
				Review:   "Profesional dan cepat dalam merapikan taman di rumah saya. Akan memanggilnya lagi jika ada masalah lain.",
			},
			{
				WorkerID: "5d4bc029-2d99-4fa2-b853-e24febafee1d",
				UserID:   "51487321-5387-4204-a788-686f7b1b80d4",
				Rating:   4.2,
				Review:   "Sangat ahli dalam memperbaiki komputer di rumah saya. Pasti akan memanggilnya lagi jika ada masalah lain.",
			},
			{
				WorkerID: "5d4bc029-2d99-4fa2-b853-e24febafee1d",
				UserID:   "d296f037-e55d-400f-8fb6-21dc97cc8fad",
				Rating:   4.0,
				Review:   "Cepat dan profesional dalam memperbaiki bug yang ada dalam komputer. Akan memanggilnya lagi jika ada masalah lain.",
			},
		}

		if err := db.Model(&domain.Review{}).CreateInBatches(&reviews, len(reviews)).Error; err != nil {
			return err
		}

		return nil
	}
}
