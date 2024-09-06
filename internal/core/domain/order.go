package domain

import (
	"time"

	"github.com/lib/pq"
)

type Order struct {
	OrderID       string        `gorm:"type:varchar(36)"`
	Date          time.Time     `gorm:"type:timestampz"`
	Time          time.Time     `gorm:"type:timestampz"`
	WorkerService pq.Int64Array `gorm:"type:integer[]"`
	Address       string        `gorm:"type:text"`
	Type          string        `gorm:"type:text"`
	DetailAddress string        `gorm:"type:text"`
	PaymentType   string        `gorm:"type:varchar(255)"`
	PromoCode     string        `gorm:"type:varchar(255)"`
	ServiceFee    float64       `gorm:"type:decimal"`
	TotalPrice    float64       `gorm:"type:decimal"`
	OrderStatus   OrderStatus   `gorm:"type:integer"`
	CreatedAt     time.Time     `gorm:"timestampz;autoCreateTime"`
	UpdateAt      time.Time     `gorm:"timestampz;autoUpdateTime"`
}

type OrderStatus uint64

var (
	OnProcess OrderStatus = 1
	Finished  OrderStatus = 2
	Canceled  OrderStatus = 3
)
