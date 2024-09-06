package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	FullName  string    `gorm:"type:varchar(255)"`
	Phone     string    `gorm:"type:varchar(15);unique"`
	Email     string    `gorm:"type:varchar(255);unique"`
	Password  string    `gorm:"type:varchar(255)"`
	Gender    string    `gorm:"type:varchar(255)"`
	Birth     time.Time `gorm:"type:timestamp"`
	IsActive  bool      `gorm:"type:boolean;default:false"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoUpdateTime"`
}
