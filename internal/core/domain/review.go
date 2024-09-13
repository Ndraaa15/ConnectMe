package domain

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID          uint64    `gorm:"serial,primaryKey"`
	UserID      uuid.UUID `gorm:"type:varchar(36)"`
	User        User      `gorm:"references:ID;foreignKey:UserID"`
	WorkerID    uuid.UUID `gorm:"type:varchar(36)"`
	Rating      float64   `gorm:"type:decimal(2,1)"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"timestamp;autoCreateTime"`
	UpdateAt    time.Time `gorm:"timestamp;autoUpdateTime"`
}
