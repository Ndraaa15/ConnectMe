package domain

import "time"

type Bot struct {
	ID        uint64    `gorm:"serial,primaryKey"`
	Problem   string    `gorm:"type:text"`
	Picture   string    `gorm:"type:text"`
	Response  string    `gorm:"type:text"`
	Keyword   []string  `gorm:"type:text[]"`
	CreatedAt time.Time `gorm:"timestampz;autoCreateTime"`
	UpdatedAt time.Time `gorm:"timestampz;autoUpdateTime"`
}
