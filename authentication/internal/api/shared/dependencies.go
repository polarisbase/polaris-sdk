package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users"
)

type Dependencies struct {
	FiberRouter         fiber.Router
	UserActionsProvider *users.ActionsProvider
}

func NewDependencies(fiberRouter fiber.Router, userActionsProvider *users.ActionsProvider) *Dependencies {

	d := &Dependencies{}

	d.FiberRouter = fiberRouter
	d.UserActionsProvider = userActionsProvider

	return d

}
