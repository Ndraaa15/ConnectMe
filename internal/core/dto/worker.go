package dto

type WorkerResponse struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Tag        TagResponse    `json:"tag"`
	LowerPrice float64        `json:"lower_price"`
	Image      string         `json:"image"`
	Review     ReviewResponse `json:"review"`
}

type WorkerDetailResponse struct {
	ID             string                  `json:"id"`
	Name           string                  `json:"name"`
	Tag            TagResponse             `json:"tag"`
	Description    string                  `json:"description"`
	WorkExperience uint64                  `json:"work_experience"`
	LowerPrice     float64                 `json:"lower_price"`
	WorkerServices []WorkerServiceResponse `json:"worker_services"`
	WorkHour       []string                `json:"work_hour"`
	Image          string                  `json:"image"`
	Review         ReviewResponse          `json:"review"`
}

type WorkerServiceResponse struct {
	ID      uint64  `json:"id"`
	Service string  `json:"service"`
	Price   float64 `json:"price"`
}

type GetWorkersFilter struct {
	Keyword     string
	IsPopular   bool
	LowerPrice  float64
	HigherPrice float64
}
