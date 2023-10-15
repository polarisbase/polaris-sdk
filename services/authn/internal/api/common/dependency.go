package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/info"
)

type Dependencies struct {
	FiberRouter fiber.Router
	InfoActions *info.ActionsProvider
}

func NewDependencies(fiberRouter fiber.Router, infoActionsProvider *info.ActionsProvider) *Dependencies {

	d := &Dependencies{}

	d.FiberRouter = fiberRouter

	d.InfoActions = infoActionsProvider

	return d

}
