package v1

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/api/shared"
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

	// Logout
	a.dep.FiberRouter.Post(
		fmt.Sprintf("%s/logout", a.prefix),
		a.Logout,
	)

	// Logout (GET)
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/logout", a.prefix),
		a.Logout,
	)

	// User info
	a.dep.FiberRouter.Get(
		fmt.Sprintf("%s/user-info", a.prefix),
		a.UserInfo,
	)

	return a

}
