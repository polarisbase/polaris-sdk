package sessions

import (
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/persistence"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/sessions/shared"
	v1 "github.com/polarisbase/polaris-sdk/v2/authentication/internal/sessions/v1"
)

type ActionsProvider struct {
	dep *shared.Dependencies
	V1  *v1.Actions
}

func NewActionsProvider(persist persistence.SqlLite) *ActionsProvider {
	// Create the actions provider
	ap := &ActionsProvider{}
	// Create the shared dependencies
	ap.dep = shared.NewDependencies(persist)
	// Create the v1 actions
	ap.V1 = v1.New(ap.dep)
	// Return the actions provider
	return ap
}
