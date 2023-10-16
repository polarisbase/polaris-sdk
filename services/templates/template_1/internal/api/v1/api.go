package v1

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/v3/services/templates/template_1/internal/api/common"
)

type Api struct {
	prefix string
	dep    *common.Dependencies
}

func New(dependencies *common.Dependencies) *Api {

	a := &Api{}

	a.prefix = "v1"

	a.dep = dependencies

	// Create an info item
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/_/create-info", a.prefix),
		a.CreateInfo,
	)

	// List the info items
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/_/list-info", a.prefix),
		a.ListInfo,
	)

	return a

}
