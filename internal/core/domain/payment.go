package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	PaymentID         uuid.UUID     `gorm:"varchar(36);primaryKey"`
	OrderID           uuid.UUID     `gorm:"type:varchar(36)"`
	ServiceFee        float64       `gorm:"type:decimal"`
	TotalServicePrice float64       `gorm:"type:decimal"`
	TotalPrice        float64       `gorm:"type:decimal"`
	PaymentType       PaymentType   `gorm:"type:integer"`
	PromoCode         string        `gorm:"type:varchar(255)"`
	Status            StatusPayment `gorm:"type:integer"`
	CreatedAt         time.Time     `gorm:"timestamp;autoCreateTime"`
	UpdatedAt         time.Time     `gorm:"timestamp;autoUpdateTime"`
}

type PaymentType uint64

const (
	PaymentTypeUnknown   PaymentType = 0
	PaymentTypeBCAVa     PaymentType = 1
	PaymentTypeBRIVa     PaymentType = 2
	PaymentTypeMandiriVa PaymentType = 3
	PaymentTypeBNIVa     PaymentType = 4
	PaymentTypePermataVa PaymentType = 5
	PaymentTypeGopay     PaymentType = 6
	PaymentTypeShopeePay PaymentType = 7
)

var PaymentTypeMap = map[PaymentType]string{
	PaymentTypeBCAVa:     "BCA Virtual Account",
	PaymentTypeBRIVa:     "BRI Virtual Account",
	PaymentTypeMandiriVa: "Mandiri Virtual Account",
	PaymentTypeBNIVa:     "BNI Virtual Account",
	PaymentTypePermataVa: "Permata Virtual Account",
	PaymentTypeGopay:     "Gopay",
	PaymentTypeShopeePay: "ShopeePay",
}

func (s PaymentType) String() string {
	return PaymentTypeMap[s]
}

func (s PaymentType) Value() uint64 {
	return uint64(s)
}

type StatusPayment uint64

const (
	StatusPaymentUnknown StatusPayment = 0
	StatusPaymentSuccess StatusPayment = 1
	StatusPaymentOnGoing StatusPayment = 2
	StatusPaymentFailed  StatusPayment = 3
)

var StatusPaymentMap = map[StatusPayment]string{
	StatusPaymentSuccess: "Success",
	StatusPaymentOnGoing: "On Going",
	StatusPaymentFailed:  "Failed",
}

func (s StatusPayment) String() string {
	return StatusPaymentMap[s]
}

func (s StatusPayment) Value() uint64 {
	return uint64(s)
}
