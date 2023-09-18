package component

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	_ "github.com/gofiber/fiber/v2/middleware/adaptor"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"net/http"
)

func CreateHandler(title string, body g.Node) fiber.Handler {
	return adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Rendering a Node is as simple as calling Render and passing an io.Writer
		_ = Page(title, r.URL.Path, body).Render(w)
	})
}

func Page(title, path string, body g.Node) g.Node {
	// HTML5 boilerplate document
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=typography")),
		},
		Body: []g.Node{
			Navbar(path, []PageLink{
				{Path: "/contact", Name: "Contact"},
				{Path: "/about", Name: "About"},
			}),
			Container(
				Prose(body),
				PageFooter(),
			),
		},
	})
}
