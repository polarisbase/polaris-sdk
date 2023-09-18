package common

import (
	"github.com/gofiber/fiber/v2"
)

type Site interface {
	EntryPoint() fiber.Handler
}
