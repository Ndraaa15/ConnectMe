package domain

import (
	"time"
)

type Review struct {
	UserID    string    `gorm:"type:varchar(36);primaryKey"`
	User      User      `gorm:"references:ID;foreignKey:UserID"`
	WorkerID  string    `gorm:"type:varchar(36);primaryKey"`
	Rating    float64   `gorm:"type:decimal(2,1)"`
	Review    string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"timestamp;autoCreateTime"`
	UpdateAt  time.Time `gorm:"timestamp;autoUpdateTime"`
}
