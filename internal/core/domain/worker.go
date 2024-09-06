package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Worker struct {
	ID             uuid.UUID       `gorm:"type:varchar(36);primaryKey"`
	TagID          uint64          `gorm:"type:integer"`
	Tag            Tag             `gorm:"references:ID;foreignKey:TagID"`
	Desxription    string          `gorm:"type:text"`
	WorkExperience uint64          `gorm:"type:integer"`
	LowerPrice     float64         `gorm:"type:decimal"`
	WorkerServices []WorkerService `gorm:"references:ID;foreignKey:WorkerID;constraint:OnDelete:CASCADE"`
	WorkHour       pq.StringArray  `gorm:"type:varchar(5)[]"`
	CreatedAt      time.Time       `gorm:"type:timestampz;autoCreateTime"`
	UpdatedAt      time.Time       `gorm:"type:timestampz;autoUpdateTime"`
}

type WorkerService struct {
	ID        uint64    `gorm:"type:serial;primaryKey"`
	WorkerID  uuid.UUID `gorm:"type:varchar(36)"`
	Service   string    `gorm:"type:varchar(255)"`
	Price     float64   `gorm:"type:decimal"`
	CreatedAt time.Time `gorm:"type:timestampz;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestampz;autoUpdateTime"`
}
