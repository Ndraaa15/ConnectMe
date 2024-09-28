package seed

import (
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/bcrypt"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"gorm.io/gorm"
)

func UserSeeder() Seeder {
	return func(db *gorm.DB) error {
		hashedPassword, err := bcrypt.EncryptPassword("password")
		if err != nil {
			return err
		}

		birth, err := time.Parse("2006-01-02", "2003-12-15")
		if err != nil {
			return err
		}

		users := []domain.User{
			{
				ID:       "3eae758c-7c50-4266-b28d-eb180ffd8336",
				FullName: "Gede Indra Adi Brata",
				Email:    "indrabrata599@gmail.com",
				Phone:    "+628123456789",
				Password: hashedPassword,
				Gender:   domain.GenderMale,
				Birth:    birth,
				Role:     domain.RoleUser,
				IsActive: true,
			},
			{
				ID:       "d296f037-e55d-400f-8fb6-21dc97cc8fad",
				FullName: "Paula Sugiarto",
				Email:    "paulaaaa@gmail.com",
				Phone:    "+628123456111",
				Password: hashedPassword,
				Gender:   domain.GenderFemale,
				Birth:    birth,
				Role:     domain.RoleUser,
				IsActive: true,
			},
			{
				ID:       "5d4bc029-2d99-4fa2-b853-e24febafee1d",
				FullName: "Handedius Sando Sianipar",
				Email:    "sandogi@gmail.com",
				Phone:    "+628123456112",
				Password: hashedPassword,
				Gender:   domain.GenderMale,
				Birth:    birth,
				Role:     domain.RoleWorker,
				IsActive: true,
			},
		}

		if err := db.Model(&domain.User{}).CreateInBatches(&users, len(users)).Error; err != nil {
			return err
		}

		return nil
	}
}
