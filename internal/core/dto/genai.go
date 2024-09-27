package dto

type ResponseProblem struct {
	Solution string   `json:"solution"`
	Keyword  []string `json:"keyword"`
}
