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

	a.prefix = "v1"

	a.dep = dependencies

	// Create the user API
	a.dep.FiberRouter.Post(
		fmt.Sprintf("%s/users", a.prefix),
		a.Create,
	)

	// List the users
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/users", a.prefix),
		a.List,
	)

	// Login
	a.dep.FiberRouter.Post(
		fmt.Sprintf("%s/login", a.prefix),
		a.Login,
	)

	return a

}
