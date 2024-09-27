package dto

type BotResponse struct {
	Problem  string           `json:"problem"`
	Image    string           `json:"image"`
	Solution string           `json:"solution"`
	Workers  []WorkerResponse `json:"workers"`
}
