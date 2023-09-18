package webui

import (
	"fmt"
	"net/url"

	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webui/common"
)

type WebUI struct {
	prefix string
	api *api.Service
	pages []common.Page
}

// RegisterApi registers the api service.
func (ui *WebUI) Bootstrap(api *api.Service) {
	ui.api = api
	for _, page := range ui.pages {
		path,_ := url.JoinPath(ui.prefix,page.GetPath())
		fmt.Printf("Registering page %s\n", path)
		api.DirectAccessFiberRouter().Get(
			path, 
			page.GetFiberHandler())
	}
}

// AddPage adds a page.
func (ui *WebUI) AddPage(page common.Page) {
	found := false
	for _, p := range ui.pages {
		if p.GetPath() == page.GetPath() {
			// TODO: Log error
			found = true
			break
		}
	}

	if !found {
		ui.pages = append(ui.pages, page)
	}
}

// New creates a new WebUI.
func New(prefix string) *WebUI {
	ui := &WebUI{
		prefix: prefix,
		pages: []common.Page{},
	}

	return ui
}