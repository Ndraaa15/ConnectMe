package domain

import "time"

type Favourite struct {
	UserID    string    `gorm:"type:varchar(36);primaryKey"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	WorkerID  string    `gorm:"type:varchar(36);primaryKey"`
	Worker    Worker    `gorm:"foreignKey:WorkerID;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoUpdateTime"`
}
