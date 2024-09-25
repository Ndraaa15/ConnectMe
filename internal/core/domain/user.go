package domain

import (
	"time"
)

type User struct {
	ID        string      `gorm:"type:varchar(36);primaryKey"`
	FullName  string      `gorm:"type:varchar(255)"`
	Phone     string      `gorm:"type:varchar(15);unique"`
	Email     string      `gorm:"type:varchar(255);unique"`
	Password  string      `gorm:"type:varchar(255)"`
	Gender    Gender      `gorm:"type:integer"`
	Birth     time.Time   `gorm:"type:timestamp"`
	IsActive  bool        `gorm:"type:boolean;default:false"`
	Role      AccountRole `gorm:"type:integer"`
	CreatedAt time.Time   `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time   `gorm:"type:timestamp;autoUpdateTime"`
}

type Gender uint64

const (
	GenderUnknown Gender = 0
	GenderMale    Gender = 1
	GenderFemale  Gender = 2
)

var GenderMap = map[Gender]string{
	GenderMale:   "Male",
	GenderFemale: "Female",
}

func (s Gender) String() string {
	return GenderMap[s]
}

func (s Gender) Value() uint64 {
	return uint64(s)
}

type AccountRole uint64

const (
	RoleUnknown AccountRole = 0
	RoleUser    AccountRole = 1
	RoleWorker  AccountRole = 2
)

var AccountRoleMap = map[AccountRole]string{
	RoleUser:   "user",
	RoleWorker: "worker",
}

func (s AccountRole) String() string {
	return AccountRoleMap[s]
}

func (s AccountRole) Value() uint64 {
	return uint64(s)
}
