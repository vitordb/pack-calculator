package ports

type Result struct {
	ID        string      `json:"id"`
	Amount    int         `json:"amount"`
	PackSizes []int       `json:"pack_sizes"`
	Solution  map[int]int `json:"solution"`
}

type DBInterface interface {
	SaveResult(result *Result) error
	GetResultByID(id string) (*Result, error)
}
