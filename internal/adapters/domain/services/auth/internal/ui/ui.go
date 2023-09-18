package ui

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webui"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webui/common"
)

type WebUI struct {
	common.WebUI
}

func New(prefix string) *WebUI {

	ui := &WebUI{
		WebUI: webui.New(prefix),
	}

	ui.AddIndex()

	return ui
	
}
