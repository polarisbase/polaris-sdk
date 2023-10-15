package v1

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/api/common"
)

type Api struct {
	prefix string
	dep    *common.Dependencies
}

func New(dependencies *common.Dependencies) *Api {

	a := &Api{}

	a.prefix = "v1"

	a.dep = dependencies

	// Create the user API
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s", a.prefix),
		a.Info,
	)

	return a

}
