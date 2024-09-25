package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Worker struct {
	ID             uuid.UUID       `gorm:"type:varchar(36);primaryKey"`
	Name           string          `gorm:"type:varchar(255)"`
	TagID          uint64          `gorm:"type:integer"`
	Tag            Tag             `gorm:"references:ID;foreignKey:TagID"`
	Description    string          `gorm:"type:text"`
	WorkExperience uint64          `gorm:"type:integer"`
	LowerPrice     float64         `gorm:"-"`
	WorkerServices []WorkerService `gorm:"references:ID;foreignKey:WorkerID;constraint:OnDelete:CASCADE"`
	Image          string          `gorm:"type:text"`
	WorkHour       pq.StringArray  `gorm:"type:varchar(5)[]"`
	Rating         float64         `gorm:"-"`
	TotalRating    uint64          `gorm:"-"`
	TotalReview    uint64          `gorm:"-"`
	Reviews        []Review        `gorm:"references:ID;foreignKey:WorkerID;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time       `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt      time.Time       `gorm:"type:timestamp;autoUpdateTime"`
}

type WorkerService struct {
	ID        uint64    `gorm:"type:serial;primaryKey"`
	WorkerID  uuid.UUID `gorm:"type:varchar(36)"`
	Service   string    `gorm:"type:varchar(255)"`
	Price     float64   `gorm:"type:decimal"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoUpdateTime"`
}
