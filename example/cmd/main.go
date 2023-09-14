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
		pbsdk.Options.ApiService.SetFiberPortsToListen([]string{":8080"}),
		pbsdk.Options.PostgresService.SetPostgresConnection("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		
	)

	// Start the application
	app.Api.Start()
}