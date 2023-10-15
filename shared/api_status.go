package shared

type ApiResponseStatus string

type ApiMetadata struct {
	StatusCode int               `json:"status_code"`
	Status     ApiResponseStatus `json:"status"`
}
