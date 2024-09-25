package dto

import "github.com/Ndraaa15/ConnectMe/internal/core/domain"

type CreateOrderRequest struct {
	WorkerID      string              `json:"workerID" validate:"required"`
	WorkerService []int64             `json:"workerService" validate:"required"`
	OrderAddress  OrderAddressRequest `json:"orderAddress" validate:"required"`
	Date          string              `json:"date" validate:"required"`
	Time          string              `json:"time" validate:"required"`
	Payment       PaymentRequest      `json:"payment" validate:"required"`
}

type OrderAddressRequest struct {
	Street        string  `json:"street" validate:"required"`
	Latitude      float64 `json:"latitude" validate:"required"`
	Longitude     float64 `json:"longitude" validate:"required"`
	AddressType   string  `json:"addressType" validate:"required"`
	DetailAddress string  `json:"detailAddress"`
}

type UpdateOrderRequest struct {
	Status string `json:"status"`
}

type OrderResponse struct {
	OrderID     string      `json:"orderID"`
	StatusOrder string      `json:"statusOrder"`
	WorkerID    string      `json:"workerID"`
	WorkerName  string      `json:"workerName"`
	WorkerImage string      `json:"workerImage"`
	Tag         TagResponse `json:"tag"`
	TotalPrice  float64     `json:"totalPrice"`
	OrderDate   string      `json:"orderDate"`
	OrderTime   string      `json:"orderTime"`
}

type OrderDetailResponse struct {
	OrderID         string                  `json:"orderID"`
	StatusOrder     string                  `json:"statusOrder"`
	WorkerID        string                  `json:"workerID"`
	WorkerName      string                  `json:"workerName"`
	WorkerImage     string                  `json:"workerImage"`
	Tag             TagResponse             `json:"tag"`
	TransactionTime string                  `json:"transactionTime"`
	PaymentMethod   string                  `json:"paymentMethod"`
	PaymentStatus   string                  `json:"paymentStatus"`
	Location        string                  `json:"location"`
	WorkerService   []WorkerServiceResponse `json:"workerService"`
	ServiceFee      float64                 `json:"serviceFee"`
}

type GetOrderFilter struct {
	Status []domain.StatusOrder
}
