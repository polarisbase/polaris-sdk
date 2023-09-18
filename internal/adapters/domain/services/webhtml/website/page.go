package website

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	c "github.com/maragudk/gomponents/components"
	"net/http"
)

type BasePage struct {
	Path string
	c.HTML5Props
}

func (p *BasePage) GetPath() string {
	return p.Path
}

func (p *BasePage) GetTitle() string {
	return p.Title
}

func (p *BasePage) GetHandler() fiber.Handler {
	return adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Rendering a Node is as simple as calling Render and passing an io.Writer
		_ = c.HTML5(p.HTML5Props).Render(w)
	})
}

func NewBasePage(path string, title string) *BasePage {
	p := &BasePage{
		Path: path,
	}

	p.Title = title

	return p
}
