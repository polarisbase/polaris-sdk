package postgres

import "github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"

type Option struct {
	*common.BaseOption
}

func NewOption(optionName string, instanceName string, optionFunction func(obj interface{}) error) *Option {
	return &Option{
		BaseOption: &common.BaseOption{
			ServiceInstanceName: instanceName,
			ServiceName:         common.DB_POSTGRES_SERVICE,
			OptionName:          optionName,
			OptionFunction:      optionFunction,
		},
	}
}
