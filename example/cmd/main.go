package main

import (
	"example"

	pbsdk "github.com/polarisbase/polaris-sdk"
)

func main() {
	println("Hello World!")

	// Create a new application
	app := example.NewApp(
		// Set the ports to listen
		pbsdk.Options.ApiService.SetFiberPortsToListen("main-api", []string{":8080"}),
		pbsdk.Options.PostgresService.SetPostgresConnection("main-api", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
	)

	// Test register new user
	if userId, err := app.Auth.GetActions().RegisterNewUser("me@travishills.com", "password123"); err != nil {
		panic(err)
	} else {
		println("New user registered with id:", userId)
	}

	// Start the application
	app.Api.Start()
}
