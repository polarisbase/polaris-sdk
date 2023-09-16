package example

import (
	pbsdk "github.com/polarisbase/polaris-sdk"
	"github.com/polarisbase/polaris-sdk/pbshared"
)

type App struct {
	Api  pbshared.ApiService
	Auth pbshared.AuthService
}

func NewApp(options ...pbshared.Option) *App {
	app := &App{}

	// Create a new API service carl
	app.Api = pbsdk.NewApiService(
		"main-api",
		options...,
	)

	// Create a new Auth service
	app.Auth = pbsdk.NewAuthService(
		"main-auth",
		pbsdk.Options.ApiService.SetFiberApp("main-auth", app.Api.DirectAccessFiber()),
		pbsdk.Options.ApiService.UseFiberRouterPrefix("main-auth", "/auth"),
	)

	return app
}
