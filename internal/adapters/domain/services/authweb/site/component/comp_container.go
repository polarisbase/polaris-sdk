package component

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Container(children ...g.Node) g.Node {
	return Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}

func Prose(children ...g.Node) g.Node {
	return Div(Class("prose"), g.Group(children))
}
