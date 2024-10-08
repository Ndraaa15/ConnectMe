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
				Tag:            "Komputer",
				Specialization: "Komputer/Perangkat Keras",
			},
			{
				ID:             2,
				Tag:            "Kebun",
				Specialization: "Kebun/Taman",
			},
			{
				ID:             3,
				Tag:            "Desainer Grafis",
				Specialization: "Desain 2d/3d",
			},
			{
				ID:             4,
				Tag:            "Penulis",
				Specialization: "Penulisan Konten",
			},
			{
				ID:             5,
				Tag:            "Fotografer",
				Specialization: "Fotografi Pernikahan",
			},
			{
				ID:             6,
				Tag:            "Pengembang Web",
				Specialization: "Pengembangan Web",
			},
			{
				ID:             7,
				Tag:            "Penerjemah",
				Specialization: "Penerjemahan",
			},
			{
				ID:             8,
				Tag:            "Pengelola Media Sosial",
				Specialization: "Media Sosial",
			},
			{
				ID:             9,
				Tag:            "Videografer",
				Specialization: "Videografi",
			},
		}

		if err := db.Model(&domain.Tag{}).CreateInBatches(&tags, len(tags)).Error; err != nil {
			return err
		}

		return nil
	}
}
