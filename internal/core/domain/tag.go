package domain

import "time"

type Tag struct {
	ID        uint64    `gorm:"serial,primaryKey"`
	Tag       string    `gorm:"varchar(255),unique"`
	CreatedAt time.Time `json:"-" gorm:"timestamp;autoCreateTime"`
	UpdateAt  time.Time `json:"-" gorm:"timestamp;autoUpdateTime"`
}
