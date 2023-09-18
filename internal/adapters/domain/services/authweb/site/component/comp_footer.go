package component

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"time"
)

func PageFooter() g.Node {
	return Footer(Class("prose prose-sm prose-indigo"),
		P(
			// We can use string interpolation directly, like fmt.Sprintf.
			g.Textf("Rendered %v. ", time.Now().Format(time.RFC3339)),

			// Conditional inclusion
			g.If(time.Now().Second()%2 == 0, g.Text("It's an even second.")),
			g.If(time.Now().Second()%2 == 1, g.Text("It's an odd second.")),
		),

		P(A(Href("https://www.gomponents.com"), g.Text("gomponents"))),
	)
}
