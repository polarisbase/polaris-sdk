package v1

import "github.com/polarisbase/polaris-sdk/v2/authentication/internal/sessions/shared"

type Actions struct {
	dep *shared.Dependencies
}

func New(dependencies *shared.Dependencies) *Actions {

	a := &Actions{}

	a.dep = dependencies

	return a

}
