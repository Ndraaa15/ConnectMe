package migration

import (
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"gorm.io/gorm"
)

func MigrateUp(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&domain.User{},
	); err != nil {
		return err
	}

	return nil
}

func MigrateDown(db *gorm.DB) error {
	if err := db.Migrator().DropTable(
		&domain.User{},
	); err != nil {
		return err
	}

	return nil
}
