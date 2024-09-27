package domain

import "time"

type Bot struct {
	ID        uint64    `gorm:"serial,primaryKey"`
	UserID    string    `gorm:"type:varchar(255)"`
	Problem   string    `gorm:"type:text"`
	Picture   string    `gorm:"type:text"`
	Solution  string    `gorm:"type:text"`
	Keyword   []string  `gorm:"type:text[]"`
	CreatedAt time.Time `gorm:"timestamp;autoCreateTime"`
	UpdatedAt time.Time `gorm:"timestamp;autoUpdateTime"`
}
