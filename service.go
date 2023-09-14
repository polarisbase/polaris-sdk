package pbsdk

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/postgres"
	"github.com/polarisbase/polaris-sdk/pbshared"
)


// NewApiService creates a new API service
func NewApiService(applicationName string, options ...pbshared.Option) pbshared.ApiService {
	api_ := api.New(applicationName, options...)
	return api_
}

func NewPostgresService(applicationName string, options ...pbshared.Option) pbshared.PostgresService {
	postgres := postgres.New(applicationName, options...)
	return postgres
}
