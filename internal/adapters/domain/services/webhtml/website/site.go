package website

import (
	"github.com/gofiber/fiber/v2"
	webCommon "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webhtml/common"
)

type BaseSite struct {
	Pages map[string]webCommon.Page
}

func (s *BaseSite) GetPages() map[string]webCommon.Page {
	return s.Pages
}

func (s *BaseSite) AddPage(page webCommon.Page) {
	s.Pages[page.GetPath()] = page
}

func (s *BaseSite) EntryPoint(router fiber.Router) {
	for _, page := range s.Pages {
		page.EntryPoint(router)
	}
}

func NewBaseSite() *BaseSite {
	return &BaseSite{
		Pages: make(map[string]webCommon.Page),
	}
}
