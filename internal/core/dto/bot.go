package dto

type BotResponse struct {
	Problem  string
	Image    string
	Solution string
	Worker   []WorkerResponse
}
