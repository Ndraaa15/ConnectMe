package domain

import "time"

type PromoCode struct {
	ID        uint64    `gorm:"serial;primaryKey"`
	Code      string    `gorm:"type:varchar(255)"`
	Discount  uint64    `gorm:"type:integer"`
	IsActive  bool      `gorm:"type:boolean;default:true"`
	CreatedAt time.Time `gorm:"timestamp;autoCreateTime"`
	UpdatedAt time.Time `gorm:"timestamp;autoUpdateTime"`
}
