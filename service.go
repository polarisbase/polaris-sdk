package pbsdk

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/postgres"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webui"
	"github.com/polarisbase/polaris-sdk/pbshared"
)

// NewApiService creates a new API service
func NewApiService(applicationName string, options ...pbshared.Option) pbshared.ApiService {
	api_ := api.New(applicationName, options...)
	return api_
}

// NewPostgresService creates a new Postgres service
func NewPostgresService(applicationName string, options ...pbshared.Option) pbshared.PostgresService {
	postgres := postgres.New(applicationName, options...)
	return postgres
}

// NewAuthService creates a new Auth services
func NewAuthService(applicationName string, options ...pbshared.Option) pbshared.AuthService {
	authService := auth.New(applicationName, options...)
	return authService
}

// NewWebUI creates a new WebUI
func NewWebUI(sitePrefix string) pbshared.WebUI {
	webUI := webui.New(sitePrefix)
	return webUI
}
