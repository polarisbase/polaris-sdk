package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
	userCommon "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"
)

type Api struct {
	Api     *api.Service
	Actions userCommon.Actions
}

func New(
	api *api.Service,
	actions userCommon.Actions,
) *Api {
	a := &Api{}

	// Create the service.
	a.Api = api

	// Create the actions.
	a.Actions = actions

	// Register routes.
	a.RegisterRoutes()

	return a
}

func (a *Api) RegisterRoutes() {
	a.Api.DirectAccessFiberRouter().Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
