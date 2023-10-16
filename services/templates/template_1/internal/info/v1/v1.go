package v1

import "github.com/polarisbase/polaris-sdk/v3/services/templates/template_1/internal/info/common"

type Actions struct {
	dep *common.Dependencies
}

func New(dependencies *common.Dependencies) *Actions {

	a := &Actions{}

	a.dep = dependencies

	return a

}
