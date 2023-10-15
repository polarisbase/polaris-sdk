package contracts

type LoginResponse struct {
	AccessCode string `json:"access_code"`
	ExpiresIn  int64  `json:"expires_in"`
}
