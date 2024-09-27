package seed

import (
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func WorkerSeeder() Seeder {
	return func(db *gorm.DB) error {
		workers := []domain.Worker{
			{
				ID:             "5d4bc029-2d99-4fa2-b853-e24febafee1d",
				Name:           "Indra",
				TagID:          1,
				Description:    "Teknisi komputer yang sudah memiliki 5 tahun pengalaman.",
				WorkExperience: 4,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Permbersihan Virus",
						Price:    200000,
						WorkerID: "5d4bc029-2d99-4fa2-b853-e24febafee1d",
					},
					{
						Service:  "Pembersihan Malware",
						Price:    250000,
						WorkerID: "5d4bc029-2d99-4fa2-b853-e24febafee1d",
					},
				},
				Image:    "https://example.com/images/desain-logo.jpg",
				WorkHour: pq.StringArray{"09:00", "12:00", "14:00", "17:00"},
			},
			{
				ID:             "69f74995-fd37-4f32-8583-e8baad2cd18f",
				Name:           "Sando",
				TagID:          3,
				Description:    "Tukang Rumput yang sudah memiliki 5 tahun pengalaman.",
				WorkExperience: 4,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Perawatan Rumput",
						Price:    200000,
						WorkerID: "69f74995-fd37-4f32-8583-e8baad2cd18f",
					},
					{
						Service:  "Pembuatan Taman",
						Price:    250000,
						WorkerID: "69f74995-fd37-4f32-8583-e8baad2cd18f",
					},
				},
				Image:    "https://example.com/images/desain-logo.jpg",
				WorkHour: pq.StringArray{"09:00", "12:00", "14:00", "17:00"},
			},
			{
				ID:             "23a60b90-9bfa-4a0a-bbe5-d454b098e437",
				Name:           "Ahmad Fauzi",
				TagID:          3,
				Description:    "Desainer grafis dengan pengalaman lebih dari 4 tahun.",
				WorkExperience: 4,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Desain Logo",
						Price:    200000,
						WorkerID: "23a60b90-9bfa-4a0a-bbe5-d454b098e437",
					},
					{
						Service:  "Desain Banner",
						Price:    250000,
						WorkerID: "23a60b90-9bfa-4a0a-bbe5-d454b098e437",
					},
				},
				Image:    "https://example.com/images/desain-logo.jpg",
				WorkHour: pq.StringArray{"09:00", "12:00", "14:00", "17:00"},
			},
			{
				ID:             "1baf6576-e77f-499b-aeef-12345b7ac5d6",
				Name:           "Budi Santoso",
				TagID:          4,
				Description:    "Penulis konten lepas dengan pengalaman lebih dari 3 tahun.",
				WorkExperience: 3,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Penulisan Artikel",
						Price:    100000,
						WorkerID: "1baf6576-e77f-499b-aeef-12345b7ac5d6",
					},
					{
						Service:  "Penulisan Blog",
						Price:    120000,
						WorkerID: "1baf6576-e77f-499b-aeef-12345b7ac5d6",
					},
				},
				Image:    "https://example.com/images/penulis-konten.jpg",
				WorkHour: pq.StringArray{"10:00", "13:00", "15:00", "18:00"},
			},
			{
				ID:             "f1a7b690-0a5b-467f-925d-0873b9e0d9f2",
				Name:           "Dian Kurniawati",
				TagID:          5,
				Description:    "Fotografer profesional dengan pengalaman lebih dari 6 tahun.",
				WorkExperience: 6,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Pemotretan Pernikahan",
						Price:    3000000,
						WorkerID: "f1a7b690-0a5b-467f-925d-0873b9e0d9f2",
					},
					{
						Service:  "Pemotretan Produk",
						Price:    1500000,
						WorkerID: "f1a7b690-0a5b-467f-925d-0873b9e0d9f2",
					},
				},
				Image:    "https://example.com/images/fotografer.jpg",
				WorkHour: pq.StringArray{"09:00", "11:00", "14:00", "16:00"},
			},
			{
				ID:             "c9b7a478-4a49-402d-8ef3-17a625f9c54e",
				Name:           "Citra Rahma",
				TagID:          6,
				Description:    "Pengembang web lepas dengan pengalaman lebih dari 5 tahun.",
				WorkExperience: 5,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Pembuatan Website",
						Price:    3500000,
						WorkerID: "c9b7a478-4a49-402d-8ef3-17a625f9c54e",
					},
					{
						Service:  "Optimisasi SEO",
						Price:    2000000,
						WorkerID: "c9b7a478-4a49-402d-8ef3-17a625f9c54e",
					},
				},
				Image:    "https://example.com/images/pengembang-web.jpg",
				WorkHour: pq.StringArray{"08:00", "11:00", "13:00", "17:00"},
			},
			{
				ID:             "05b3e0b2-9de1-44ff-9611-d402b672c65a",
				Name:           "Dedi Purwanto",
				TagID:          7,
				Description:    "Penerjemah bahasa Inggris-Indonesia dengan pengalaman lebih dari 7 tahun.",
				WorkExperience: 7,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Penerjemahan Dokumen",
						Price:    500000,
						WorkerID: "05b3e0b2-9de1-44ff-9611-d402b672c65a",
					},
					{
						Service:  "Penerjemahan Buku",
						Price:    1000000,
						WorkerID: "05b3e0b2-9de1-44ff-9611-d402b672c65a",
					},
				},
				Image:    "https://example.com/images/penerjemah.jpg",
				WorkHour: pq.StringArray{"10:00", "13:00", "15:00", "18:00"},
			},
			{
				ID:             "7b07bcb5-3b2f-4a19-9274-e87d3dbf64e8",
				Name:           "Eva Marlina",
				TagID:          8,
				Description:    "Pengelola media sosial lepas dengan pengalaman lebih dari 4 tahun.",
				WorkExperience: 4,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Manajemen Media Sosial",
						Price:    2000000,
						WorkerID: "7b07bcb5-3b2f-4a19-9274-e87d3dbf64e8",
					},
					{
						Service:  "Kampanye Media Sosial",
						Price:    3000000,
						WorkerID: "7b07bcb5-3b2f-4a19-9274-e87d3dbf64e8",
					},
				},
				Image:    "https://example.com/images/manajemen-media-sosial.jpg",
				WorkHour: pq.StringArray{"09:00", "12:00", "14:00", "17:00"},
			},
			{
				ID:             "2d93e5ea-70e5-4ecf-a8ea-f3c1c89d9f56",
				Name:           "Fajar Pratama",
				TagID:          9,
				Description:    "Videografer lepas dengan pengalaman lebih dari 5 tahun.",
				WorkExperience: 5,
				WorkerServices: []domain.WorkerService{
					{
						Service:  "Pembuatan Video Promosi",
						Price:    5000000,
						WorkerID: "2d93e5ea-70e5-4ecf-a8ea-f3c1c89d9f56",
					},
					{
						Service:  "Edit Video",
						Price:    1500000,
						WorkerID: "2d93e5ea-70e5-4ecf-a8ea-f3c1c89d9f56",
					},
				},
				Image:    "https://example.com/images/videografer.jpg",
				WorkHour: pq.StringArray{"10:00", "13:00", "16:00", "18:00"},
			},
		}

		if err := db.Model(&domain.Worker{}).CreateInBatches(&workers, len(workers)).Error; err != nil {
			return err
		}

		return nil
	}
}
