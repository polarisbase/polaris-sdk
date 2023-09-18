package pages

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webhtml/website"
)

type testPage struct {
	*website.BasePage
}

func NewTestPage() *testPage {

	p := &testPage{
		BasePage: website.NewBasePage("/test", "test page"),
	}

	p.Body = p.Render()

	return p

}

func (p *testPage) Render() []g.Node {
	return []g.Node{
		H1(nil, g.Text("Test Page")),
	}
}
