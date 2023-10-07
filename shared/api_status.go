package shared

type ApiResponseStatus string

type ApiMetadata struct {
	Status ApiResponseStatus `json:"status"`
}
