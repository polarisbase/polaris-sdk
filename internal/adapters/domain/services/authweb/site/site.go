package site

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/authweb/site/pages"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webhtml/website"
)

type Site struct {
	*website.BaseSite
}

func New() *Site {

	s := &Site{
		BaseSite: website.NewBaseSite(),
	}

	s.AddPage(pages.NewTestPage())

	return s

}
