package ui

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webui"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func (ui *WebUI) AddIndex() {

	render := func() g.Node {
		return h.Div(
			g.Text("Hello world from index!"),
		)
	}

	page := webui.NewBasePage("", render)

	ui.AddPage(page)
	
}