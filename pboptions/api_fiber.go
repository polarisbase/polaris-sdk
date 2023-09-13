package pboptions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
)

var ApiServiceOptions ApiServiceOption = ApiServiceOption{}

type ApiServiceOption struct{}

func (ApiServiceOption) SetFiberPortsToListen(portsToListen []string) *api.Option {
	return api.NewOption("set-fiber-ports-to-listen", func(obj interface{}) error {
		if service, ok := obj.(*api.Service); ok {
			service.SetFiberPortsToListen(portsToListen)
		}
		return nil
	})
}

func (ApiServiceOption) SetFiberApp(app *fiber.App) *api.Option {
	return api.NewOption("set-fiber-app", func(obj interface{}) error {
		if service, ok := obj.(*api.Service); ok {
			service.SetAsNonExacutable(true)
			service.SetFiberApp(app)
		}
		return nil
	})
}

func (ApiServiceOption) UseFiberRouterPrefix(prefix string) *api.Option {
	return api.NewOption("use-fiber-path-prefix", func(obj interface{}) error {
		if service, ok := obj.(*api.Service); ok {
			service.SetFiberRouterPrefix(prefix)
		}
		return nil
	})
}
