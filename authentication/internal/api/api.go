package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/api/shared"
	v1 "github.com/polarisbase/polaris-sdk/v2/authentication/internal/api/v1"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users"
)

type Api struct {
	dep *shared.Dependencies
	v1  *v1.Api
}

func New(fiberRouter fiber.Router, userActionsProvider *users.ActionsProvider) *Api {

	// Create the authentication API
	a := &Api{}

	// add cors middleware
	fiberRouter.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Create the shared dependencies
	a.dep = shared.NewDependencies(fiberRouter, userActionsProvider)

	// Add info route
	a.dep.FiberRouter.Get("/info", a.info)

	// Create the v1 API
	a.v1 = v1.New(a.dep)

	// Return the authentication API
	return a
}

func (a *Api) info(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"version": "v1",
	})
}
