package common

import (
	"github.com/gofiber/fiber/v2"
)

type Page interface {
	// Get the page's title.
	GetTitle() string
	GetPath() string
	GetHandler() fiber.Handler
	EntryPoint(router fiber.Router)
}
