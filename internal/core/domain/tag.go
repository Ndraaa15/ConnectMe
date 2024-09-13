package domain

import "time"

type Tag struct {
	ID             uint64    `gorm:"serial,primaryKey"`
	Tag            string    `gorm:"varchar(255),unique"`
	Specialization string    `gorm:"varchar(255)"`
	CreatedAt      time.Time `gorm:"timestamp;autoCreateTime"`
	UpdateAt       time.Time `gorm:"timestamp;autoUpdateTime"`
}
