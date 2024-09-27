package seed

import (
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func OrderSeeder() Seeder {
	return func(db *gorm.DB) error {
		orders := []domain.Order{
			{
				OrderID:       "CM-1257894000-0001",
				UserID:        "3eae758c-7c50-4266-b28d-eb180ffd8336",
				WorkerID:      "5d4bc029-2d99-4fa2-b853-e24febafee1d",
				Date:          time.Date(2024, 10, 10, 0, 0, 0, 0, time.Now().Location()),
				Time:          time.Date(0, 0, 0, 15, 30, 0, 0, time.Now().Location()),
				WorkerService: pq.Int64Array{1, 2},
				OrderStatus:   1,
			},
			{
				OrderID:       "CM-1257894001-0001",
				UserID:        "3eae758c-7c50-4266-b28d-eb180ffd8336",
				WorkerID:      "69f74995-fd37-4f32-8583-e8baad2cd18f",
				Date:          time.Date(2024, 10, 10, 0, 0, 0, 0, time.Now().Location()),
				Time:          time.Date(0, 0, 0, 15, 30, 0, 0, time.Now().Location()),
				WorkerService: pq.Int64Array{3, 4},
				OrderStatus:   2,
			},
			{
				OrderID:       "CM-1257894001-0002",
				UserID:        "3eae758c-7c50-4266-b28d-eb180ffd8336",
				WorkerID:      "69f74995-fd37-4f32-8583-e8baad2cd18f",
				Date:          time.Date(2024, 10, 10, 0, 0, 0, 0, time.Now().Location()),
				Time:          time.Date(0, 0, 0, 15, 30, 0, 0, time.Now().Location()),
				WorkerService: pq.Int64Array{5, 6},
				OrderStatus:   3,
			},
		}

		if err := db.Model(&domain.Order{}).CreateInBatches(&orders, len(orders)).Error; err != nil {
			return err
		}

		return nil
	}
}
