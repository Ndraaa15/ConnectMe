package dto

type CreateReviewRequest struct {
	WorkerID string  `json:"worker_id"`
	UserID   string  `json:"user_id"`
	Rating   float64 `json:"rating"`
	Review   string  `json:"review"`
}

type ReviewResponse struct {
	Rating        float64                `json:"rating"`
	TotalRating   uint64                 `json:"total_rating"`
	ReviewsDetail []ReviewDetailResponse `json:"reviews_detail,omitempty"`
	TotalReview   uint64                 `json:"total_review,omitempty"`
}

type ReviewDetailResponse struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Review   string  `json:"review"`
	Rating   float64 `json:"rating"`
	TimeSent string  `json:"time_sent"`
}
