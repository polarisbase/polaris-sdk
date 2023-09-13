package pbshared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
)

type Service interface {
	GetName() string
	GetServiceSystemName() common.ServiceName
	Start() error
}

type ApiService interface {
	Service
	DirectAccessFiber() *fiber.App
}

type PostgresService interface {
	Service
	Connect() error
}
