package dto

type CreateWorkerServiceRequest struct {
	WorkerID string  `json:"workerID" validate:"required"`
	Service  string  `json:"service" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type WorkerServiceResponse struct {
	ID      uint64  `json:"id"`
	Service string  `json:"service"`
	Price   float64 `json:"price"`
}
