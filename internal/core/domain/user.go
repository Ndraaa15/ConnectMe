package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	FullName  string    `gorm:"type:varchar(255)"`
	Phone     string    `gorm:"type:varchar(15);unique"`
	Email     string    `gorm:"type:varchar(255);unique"`
	Password  string    `gorm:"type:varchar(255)"`
	Gender    Gender    `gorm:"type:integer"`
	Birth     time.Time `gorm:"type:timestamp"`
	IsActive  bool      `gorm:"type:boolean;default:false"`
	CreatedAt time.Time `json:"-" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"type:timestamp;autoUpdateTime"`
}

type Gender uint64

const (
	Male   Gender = 0
	Female Gender = 1
)

var GenderMap = map[Gender]string{
	Male:   "Male",
	Female: "Female",
}

func (s Gender) String() string {
	return GenderMap[s]
}

func (s Gender) Value() uint64 {
	return uint64(s)
}
