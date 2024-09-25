package dto

type CreateOrderRequest struct {
	WorkerID      string         `json:"workerID" validate:"required"`
	WorkerService []int64        `json:"workerService" validate:"required"`
	OrderAddress  OrderAddress   `json:"orderAddress" validate:"required"`
	Date          string         `json:"date" validate:"required"`
	Time          string         `json:"time" validate:"required"`
	Payment       PaymentRequest `json:"payment" validate:"required"`
}

type OrderAddress struct {
	Street        string  `json:"street" validate:"required"`
	Latitude      float64 `json:"latitude" validate:"required"`
	Longitude     float64 `json:"longitude" validate:"required"`
	AddressType   string  `json:"addressType" validate:"required"`
	DetailAddress string  `json:"detailAddress"`
}
