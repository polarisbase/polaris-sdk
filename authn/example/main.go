package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authn"
)

func main() {
	fiberApp := fiber.New()
	authn.New(fiberApp, "api/auth")
	if err := fiberApp.Listen(":3000"); err != nil {
		panic(err)
	}
}
