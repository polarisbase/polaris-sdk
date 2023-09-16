package pbshared

import (
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetServiceInstanceName() string
	GetServiceSystemName() ServiceName
	Start() error
}

type ApiService interface {
	Service
	DirectAccessFiber() *fiber.App
	DirectAccessFiberRouter() fiber.Router
}

type PostgresService interface {
	Service
	Connect() error
}

type AuthService interface {
	Service
	GetActions() AuthActions
}
