package dto

type TagResponse struct {
	ID             uint64 `json:"id"`
	Tag            string `json:"tag"`
	Specialization string `json:"specialization"`
}
