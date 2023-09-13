package pbsdk

import (
	"github.com/polarisbase/polaris-go/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-go/internal/adapters/domain/services/api"
	"github.com/polarisbase/polaris-go/internal/adapters/domain/services/postgres"
	"github.com/polarisbase/polaris-go/pbshared"
)

func NewApiService(applicationName string, options ...common.Option) pbshared.ApiService {
	api_ := api.New(applicationName, options...)
	return api_
}

func NewPostgresService(applicationName string, options ...common.Option) pbshared.PostgresService {
	postgres := postgres.New(applicationName, options...)
	return postgres
}
