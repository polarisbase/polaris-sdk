package api

import "github.com/polarisbase/polaris-go/internal/adapters/domain/common"

type Option struct {
	*common.BaseOption
}

func NewOption(optionName string, optionFunction func(obj interface{}) error) *Option {
	return &Option{
		BaseOption: &common.BaseOption{
			ServiceName:    common.API_FIBER_SERVICE,
			OptionName:     optionName,
			OptionFunction: optionFunction,
		},
	}
}
