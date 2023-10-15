package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authn"
)

func main() {
	// Create new Fiber instance
	api := fiber.New()
	// Register authentication service
	authn.New(api, "api/auth")
	// Start server on http://localhost:3000
	if err := api.Listen(":3000"); err != nil {
		panic(err)
	}
}
