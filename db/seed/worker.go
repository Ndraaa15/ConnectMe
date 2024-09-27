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
				Description:    "Teknisi komputer berpengalaman lebih dari 5 tahun, ahli dalam perbaikan dan pemeliharaan perangkat keras serta perangkat lunak. Indra memiliki kemampuan untuk mendiagnosis dan memperbaiki berbagai masalah komputer, termasuk masalah jaringan dan keamanan. Selain itu, Indra juga berpengalaman dalam memberikan pelatihan kepada pengguna komputer.",
				WorkExperience: 5,
				WorkerServices: []domain.WorkerService{
					{
						ID:       1,
						Service:  "Pembersihan Virus",
						Price:    200000,
						WorkerID: "5d4bc029-2d99-4fa2-b853-e24febafee1d",
					},
					{
						ID:       2,
						Service:  "Pembersihan Malware",
						Price:    250000,
						WorkerID: "5d4bc029-2d99-4fa2-b853-e24febafee1d",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/teknisi%20komputer.jpg",
				WorkHour: pq.StringArray{"09:00", "12:00", "14:00", "17:00"},
			},
			{
				ID:             "69f74995-fd37-4f32-8583-e8baad2cd18f",
				Name:           "Sando",
				TagID:          3,
				Description:    "Tukang rumput dengan pengalaman lebih dari 5 tahun, ahli dalam perawatan dan pembuatan taman. Sando memiliki kemampuan untuk merancang dan merawat taman dengan berbagai jenis tanaman. Selain itu, Sando juga berpengalaman dalam penggunaan alat-alat taman modern untuk memastikan hasil yang optimal.",
				WorkExperience: 5,
				WorkerServices: []domain.WorkerService{
					{
						ID:       3,
						Service:  "Perawatan Taman",
						Price:    200000,
						WorkerID: "69f74995-fd37-4f32-8583-e8baad2cd18f",
					},
					{
						ID:       4,
						Service:  "Pembuatan Taman",
						Price:    250000,
						WorkerID: "69f74995-fd37-4f32-8583-e8baad2cd18f",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/tukang%20kebun.jpg",
				WorkHour: pq.StringArray{"09:00", "12:00", "14:00", "17:00"},
			},
			{
				ID:             "23a60b90-9bfa-4a0a-bbe5-d454b098e437",
				Name:           "Ahmad Fauzi",
				TagID:          3,
				Description:    "Desainer grafis dengan pengalaman lebih dari 4 tahun, ahli dalam desain logo dan banner. Ahmad memiliki kemampuan untuk menciptakan desain yang menarik dan profesional untuk berbagai kebutuhan bisnis. Selain itu, Ahmad juga berpengalaman dalam menggunakan berbagai perangkat lunak desain grafis modern.",
				WorkExperience: 4,
				WorkerServices: []domain.WorkerService{
					{
						ID:       5,
						Service:  "Desain Logo",
						Price:    200000,
						WorkerID: "23a60b90-9bfa-4a0a-bbe5-d454b098e437",
					},
					{
						ID:       6,
						Service:  "Desain Banner",
						Price:    250000,
						WorkerID: "23a60b90-9bfa-4a0a-bbe5-d454b098e437",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/desainer%20logo.jpg",
				WorkHour: pq.StringArray{"09:00", "12:00", "14:00", "17:00"},
			},
			{
				ID:             "1baf6576-e77f-499b-aeef-12345b7ac5d6",
				Name:           "Budi Santoso",
				TagID:          4,
				Description:    "Penulis konten lepas dengan pengalaman lebih dari 3 tahun, ahli dalam penulisan artikel dan blog. Budi memiliki kemampuan untuk menulis konten yang menarik dan informatif untuk berbagai topik. Selain itu, Budi juga berpengalaman dalam optimisasi SEO untuk meningkatkan visibilitas konten di mesin pencari.",
				WorkExperience: 3,
				WorkerServices: []domain.WorkerService{
					{
						ID:       7,
						Service:  "Penulisan Artikel",
						Price:    100000,
						WorkerID: "1baf6576-e77f-499b-aeef-12345b7ac5d6",
					},
					{
						ID:       8,
						Service:  "Penulisan Blog",
						Price:    120000,
						WorkerID: "1baf6576-e77f-499b-aeef-12345b7ac5d6",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/penulis.jpg",
				WorkHour: pq.StringArray{"10:00", "13:00", "15:00", "18:00"},
			},
			{
				ID:             "f1a7b690-0a5b-467f-925d-0873b9e0d9f2",
				Name:           "Dian Kurniawati",
				TagID:          5,
				Description:    "Fotografer profesional dengan pengalaman lebih dari 6 tahun, ahli dalam pemotretan pernikahan dan produk. Dian memiliki kemampuan untuk menangkap momen-momen indah dengan kualitas tinggi. Selain itu, Dian juga berpengalaman dalam pengeditan foto untuk memastikan hasil yang sempurna.",
				WorkExperience: 6,
				WorkerServices: []domain.WorkerService{
					{
						ID:       9,
						Service:  "Pemotretan Pernikahan",
						Price:    3000000,
						WorkerID: "f1a7b690-0a5b-467f-925d-0873b9e0d9f2",
					},
					{
						ID:       10,
						Service:  "Pemotretan Produk",
						Price:    1500000,
						WorkerID: "f1a7b690-0a5b-467f-925d-0873b9e0d9f2",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/fotograper.png",
				WorkHour: pq.StringArray{"09:00", "11:00", "14:00", "16:00"},
			},
			{
				ID:             "c9b7a478-4a49-402d-8ef3-17a625f9c54e",
				Name:           "Citra Rahma",
				TagID:          6,
				Description:    "Pengembang web lepas dengan pengalaman lebih dari 5 tahun, ahli dalam pembuatan website dan optimisasi SEO. Citra memiliki kemampuan untuk membuat website yang responsif dan user-friendly. Selain itu, Citra juga berpengalaman dalam optimisasi SEO untuk meningkatkan peringkat website di mesin pencari.",
				WorkExperience: 5,
				WorkerServices: []domain.WorkerService{
					{
						ID:       11,
						Service:  "Pembuatan Website",
						Price:    3500000,
						WorkerID: "c9b7a478-4a49-402d-8ef3-17a625f9c54e",
					},
					{
						ID:       12,
						Service:  "Optimisasi SEO",
						Price:    2000000,
						WorkerID: "c9b7a478-4a49-402d-8ef3-17a625f9c54e",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/pengembang%20web.jpg",
				WorkHour: pq.StringArray{"08:00", "11:00", "13:00", "17:00"},
			},
			{
				ID:             "05b3e0b2-9de1-44ff-9611-d402b672c65a",
				Name:           "Dedi Purwanto",
				TagID:          7,
				Description:    "Penerjemah bahasa Inggris-Indonesia dengan pengalaman lebih dari 7 tahun, ahli dalam penerjemahan dokumen dan buku. Dedi memiliki kemampuan untuk menerjemahkan berbagai jenis dokumen dengan akurasi tinggi. Selain itu, Dedi juga berpengalaman dalam penerjemahan buku untuk berbagai genre.",
				WorkExperience: 7,
				WorkerServices: []domain.WorkerService{
					{
						ID:       13,
						Service:  "Penerjemahan Dokumen",
						Price:    500000,
						WorkerID: "05b3e0b2-9de1-44ff-9611-d402b672c65a",
					},
					{
						ID:       14,
						Service:  "Penerjemahan Buku",
						Price:    1000000,
						WorkerID: "05b3e0b2-9de1-44ff-9611-d402b672c65a",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/penerjemah.jpg",
				WorkHour: pq.StringArray{"10:00", "13:00", "15:00", "18:00"},
			},
			{
				ID:             "7b07bcb5-3b2f-4a19-9274-e87d3dbf64e8",
				Name:           "Eva Marlina",
				TagID:          8,
				Description:    "Pengelola media sosial lepas dengan pengalaman lebih dari 4 tahun, ahli dalam manajemen dan kampanye media sosial. Eva memiliki kemampuan untuk mengelola akun media sosial dengan efektif dan meningkatkan engagement. Selain itu, Eva juga berpengalaman dalam merancang kampanye media sosial yang sukses.",
				WorkExperience: 4,
				WorkerServices: []domain.WorkerService{
					{
						ID:       15,
						Service:  "Manajemen Media Sosial",
						Price:    2000000,
						WorkerID: "7b07bcb5-3b2f-4a19-9274-e87d3dbf64e8",
					},
					{
						ID:       16,
						Service:  "Kampanye Media Sosial",
						Price:    3000000,
						WorkerID: "7b07bcb5-3b2f-4a19-9274-e87d3dbf64e8",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/sosial%20media%20handler.png",
				WorkHour: pq.StringArray{"09:00", "12:00", "14:00", "17:00"},
			},
			{
				ID:             "2d93e5ea-70e5-4ecf-a8ea-f3c1c89d9f56",
				Name:           "Fajar Pratama",
				TagID:          9,
				Description:    "Videografer lepas dengan pengalaman lebih dari 5 tahun, ahli dalam pembuatan video promosi dan editing video. Fajar memiliki kemampuan untuk membuat video yang menarik dan profesional untuk berbagai kebutuhan bisnis. Selain itu, Fajar juga berpengalaman dalam pengeditan video untuk memastikan hasil yang sempurna.",
				WorkExperience: 5,
				WorkerServices: []domain.WorkerService{
					{
						ID:       17,
						Service:  "Pembuatan Video Promosi",
						Price:    5000000,
						WorkerID: "2d93e5ea-70e5-4ecf-a8ea-f3c1c89d9f56",
					},
					{
						ID:       18,
						Service:  "Edit Video",
						Price:    1500000,
						WorkerID: "2d93e5ea-70e5-4ecf-a8ea-f3c1c89d9f56",
					},
				},
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/connect-me/videographer.jpg",
				WorkHour: pq.StringArray{"10:00", "13:00", "16:00", "18:00"},
			},
		}

		if err := db.Model(&domain.Worker{}).CreateInBatches(&workers, len(workers)).Error; err != nil {
			return err
		}

		return nil
	}
}
