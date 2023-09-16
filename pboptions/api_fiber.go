package pboptions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
)

var ApiServiceOptions ApiServiceOption = ApiServiceOption{}

type ApiServiceOption struct{}

func (ApiServiceOption) SetFiberPortsToListen(instanceName string, portsToListen []string) common.OptionServiceApi {
	return api.NewOption("set-fiber-ports-to-listen", instanceName, func(obj interface{}) error {
		if service, ok := obj.(*api.Service); ok {
			service.SetFiberPortsToListen(portsToListen)
		}
		return nil
	})
}

func (ApiServiceOption) SetFiberApp(instanceName string, app *fiber.App) common.OptionServiceApi {
	return api.NewOption("set-fiber-app", instanceName, func(obj interface{}) error {
		if service, ok := obj.(*api.Service); ok {
			service.SetAsNonExecutable(true)
			service.SetFiberApp(app)
		}
		return nil
	})
}

func (ApiServiceOption) UseFiberRouterPrefix(instanceName string, prefix string) common.OptionServiceApi {
	return api.NewOption("use-fiber-path-prefix", instanceName, func(obj interface{}) error {
		if service, ok := obj.(*api.Service); ok {
			service.SetFiberRouterPrefix(prefix)
		}
		return nil
	})
}
