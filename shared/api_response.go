package shared

type ApiResponse[T interface{}] struct {
	M    ApiMetadata `json:"m"`
	Data T           `json:"data"`
}
