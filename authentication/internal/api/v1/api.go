package v1

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/api/shared"
)

type Api struct {
	prefix string
	dep    *shared.Dependencies
}

func New(dependencies *shared.Dependencies) *Api {

	a := &Api{}

	a.prefix = "v1/"

	a.dep = dependencies

	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/user/test-create", a.prefix),
		a.Create,
	)

	return a

}
