package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v3/lib/persist/document/bun"
	"github.com/polarisbase/polaris-sdk/v3/services/authn"
	"github.com/polarisbase/polaris-sdk/v3/shared/fiber_middleware"
)

func main() {
	service := fiber.New()
	service.Use(fiber_middleware.Tickets())
	persistentStore := bun.New()
	if err := persistentStore.Connect(); err != nil {
		panic(err)
	}
	authn.New(service, "/api/authn", persistentStore)
	service.Listen(":5173")
}
