package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authentication"
)

func main() {
	// Create new Fiber instance
	api := fiber.New()
	// Register authentication service
	authentication.New(api, "auth")
	// Start server on http://localhost:3000
	api.Listen(":3000")
}
