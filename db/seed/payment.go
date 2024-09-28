package seed

import (
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"gorm.io/gorm"
)

func PaymentSeeder() Seeder {
	return func(db *gorm.DB) error {
		payments := []domain.Payment{
			{
				ID:                "8450ec00-e1e9-4ac0-8eb7-a7c015e7c78c",
				OrderID:           "CM-1257894000-0001",
				ServiceFee:        5000,
				TotalServicePrice: 500000,
				TotalPrice:        505000,
				PaymentType:       domain.PaymentTypeBCAVa,
				PromoCode:         "",
				Status:            domain.StatusPaymentOnGoing,
			},
			{
				ID:                "749ff038-c156-4b41-b7e7-1dc9d9c18a81",
				OrderID:           "CM-1257894001-0001",
				ServiceFee:        5000,
				TotalServicePrice: 500000,
				TotalPrice:        505000,
				PaymentType:       domain.PaymentTypeBCAVa,
				PromoCode:         "",
				Status:            domain.StatusPaymentOnGoing,
			},
			{
				ID:                "973b2d97-f4d7-45ac-bd03-2036488f6925",
				OrderID:           "CM-1257894001-0002",
				ServiceFee:        5000,
				TotalServicePrice: 500000,
				TotalPrice:        505000,
				PaymentType:       domain.PaymentTypeBCAVa,
				PromoCode:         "",
				Status:            domain.StatusPaymentOnGoing,
			},
		}

		if err := db.Model(&domain.Payment{}).CreateInBatches(&payments, len(payments)).Error; err != nil {
			return err
		}

		return nil
	}
}
