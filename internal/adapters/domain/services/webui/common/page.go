package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
)

type Page interface {
	// GetPath returns the path.
	GetPath() string
	// GetTitle returns the title.
	GetTitle() string
	// GetBody returns the body.
	GetPageHtml() g.Node
	// GetFiberHandler returns the fiber handler.
	GetFiberHandler() fiber.Handler
	// GetHandler returns the handler.
	GetHandler() http.HandlerFunc
}