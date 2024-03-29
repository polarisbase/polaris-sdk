package info

import (
	"github.com/polarisbase/polaris-sdk/v3/lib/persist"
	"github.com/polarisbase/polaris-sdk/v3/services/templates/template_1/internal/info/common"
	v1 "github.com/polarisbase/polaris-sdk/v3/services/templates/template_1/internal/info/v1"
)

type ActionsProvider struct {
	dep *common.Dependencies
	V1  *v1.Actions
}

func NewActionsProvider(persist persist.Persist) *ActionsProvider {
	// Create the actions provider
	ap := &ActionsProvider{}
	// Create the shared dependencies
	ap.dep = common.NewDependencies(persist)
	// Create the v1 actions
	ap.V1 = v1.New(ap.dep)
	// Return the actions provider
	return ap

}
