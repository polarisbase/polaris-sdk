package example

import (
	pbsdk "github.com/polarisbase/polaris-sdk"
	"github.com/polarisbase/polaris-sdk/pbshared"
)

type App struct {
	Api         pbshared.ApiService
	Auth        pbshared.AuthService
	WebHtml     pbshared.WebHtmlService
	WebAuthHtml pbshared.WebAuthHtmlService
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

	// Create a new WebHtml service
	/*app.WebHtml = pbsdk.NewWebHtmlService(
		"main-webhtml",
		pbsdk.Options.ApiService.SetFiberApp("main-webhtml", app.Api.DirectAccessFiber()),
		pbsdk.Options.ApiService.UseFiberRouterPrefix("main-webhtml", "/auth"),
	)*/

	// Create a new WebAuthHtml service
	app.WebAuthHtml = pbsdk.NewWebAuthHtmlService(
		"main-webauthhtml",
		pbsdk.Options.ApiService.SetFiberApp("main-webauthhtml", app.Api.DirectAccessFiber()),
		pbsdk.Options.ApiService.UseFiberRouterPrefix("main-webauthhtml", "/auth-site"),
	)

	return app
}
