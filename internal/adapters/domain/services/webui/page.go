package webui

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
)

type BasePage struct {
	title string
	path string
	body g.Node
	Head []g.Node
	Body []g.Node

}

// GetPath returns the path.
func (p *BasePage) GetPath() string {
	return p.path
}

// GetTitle returns the title.
func (p *BasePage) GetTitle() string {
	return p.title
}

// GetBody returns the body.
func (p *BasePage) GetPageHtml() g.Node {

	if p.Body == nil {
		p.Body = []g.Node{
			p.body,
		}
	}

	html5Props := c.HTML5Props{
		Title:    p.GetTitle(),
		Language: "en",
		Head: p.Head,
		Body: p.Body,
	}

	return c.HTML5(html5Props)
}

// GetFiberHandler returns the fiber handler.
func (p *BasePage) GetFiberHandler() fiber.Handler {
	return adaptor.HTTPHandler(p.GetHandler())
}

// GetHandler returns the handler.
func (p *BasePage) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.GetPageHtml().Render(w)
	}
}

// NewBasePage creates a new BasePage.
func NewBasePage(path string, onGetBody func () g.Node) *BasePage {
	
	bp := &BasePage{
		path: path,
		body: onGetBody(),
	}

	return bp
}