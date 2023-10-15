package contracts

type ListRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
