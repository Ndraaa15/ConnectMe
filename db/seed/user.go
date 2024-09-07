package seed

import (
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/bcrypt"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserSeeder() Seeder {
	return func(db *gorm.DB) error {
		hashedPassword, err := bcrypt.EncryptPassword("indra123")
		if err != nil {
			return err
		}

		birth, err := time.Parse("2006-01-02", "2003-12-15")
		if err != nil {
			return err
		}

		users := []domain.User{
			{
				ID:       uuid.New(),
				FullName: "Gede Indra Adi Brata",
				Email:    "indrabrata599@gmail.com",
				Phone:    "+628123456789",
				Password: hashedPassword,
				Gender:   domain.Male,
				Birth:    birth,
				IsActive: true,
			},
		}

		if err := db.CreateInBatches(&users, len(users)).Error; err != nil {
			return err
		}

		return nil
	}
}
