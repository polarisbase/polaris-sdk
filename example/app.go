package example

import (
	pbsdk "github.com/polarisbase/polaris-sdk"
	"github.com/polarisbase/polaris-sdk/pbshared"
)

type App struct {
	Api pbshared.ApiService
}

func NewApp(options ... pbshared.Option) *App {
	app := &App{}

	// Create a new API service
	app.Api = pbsdk.NewApiService(
		"example",
		options...,
	)


	return app
}