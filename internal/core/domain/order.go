package domain

import (
	"time"

	"github.com/lib/pq"
)

type Order struct {
	OrderID       string        `gorm:"type:varchar(36);primaryKey"`
	UserID        string        `gorm:"type:varchar(36)"`
	WorkerID      string        `gorm:"type:varchar(36)"`
	Worker        Worker        `gorm:"references:ID;foreignKey:WorkerID;constraint:OnDelete:CASCADE"`
	Date          time.Time     `gorm:"type:timestamp"`
	Time          time.Time     `gorm:"type:timestamp"`
	WorkerService pq.Int64Array `gorm:"type:integer[]"`
	Address       AddressOrder  `gorm:"references:OrderID;foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Payment       Payment       `gorm:"references:OrderID;foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	OrderStatus   StatusOrder   `gorm:"type:integer"`
	CreatedAt     time.Time     `gorm:"timestamp;autoCreateTime"`
	UpdateAt      time.Time     `gorm:"timestamp;autoUpdateTime"`
}

type AddressOrder struct {
	OrderID       string    `gorm:"type:varchar(255);primaryKey"`
	Street        string    `gorm:"type:text;not null"`
	Latitude      float64   `gorm:"type:decimal;not null"`
	Longitude     float64   `gorm:"type:decimal;not null"`
	AddressType   string    `gorm:"type:text;not null"`
	DetailAddress string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"timestamp;autoCreateTime"`
	UpdateAt      time.Time `gorm:"timestamp;autoUpdateTime"`
}

type StatusOrder uint64

var (
	StatusOrderUnknown  StatusOrder = 0
	StatusOrderOnGoing  StatusOrder = 1
	StatusOrderFinished StatusOrder = 2
	StatusOrderCanceled StatusOrder = 3
)

var OrderStatusMap = map[StatusOrder]string{
	StatusOrderOnGoing:  "On Going",
	StatusOrderFinished: "Finished",
	StatusOrderCanceled: "Canceled",
}

func (s StatusOrder) String() string {
	return OrderStatusMap[s]
}

func (s StatusOrder) Value() uint64 {
	return uint64(s)
}
