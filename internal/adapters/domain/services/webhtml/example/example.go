package example

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
)

type Example struct {
	apiService *api.Service
}

func NewExample(apiService *api.Service) *Example {
	e := &Example{}

	e.apiService = apiService

	e.apiService.DirectAccessFiberRouter().Get("/example",
		createHandler(e.HomePage()),
	)

	return e
}
