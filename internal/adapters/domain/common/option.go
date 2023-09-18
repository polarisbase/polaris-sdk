package common

type ServiceName string

type Option interface {
	GetServiceInstanceName() string
	GetServiceName() ServiceName
	GetOptionName() string
	ApplyOptionFunction(obj interface{}) error
}

type BaseOption struct {
	ServiceName         ServiceName
	OptionName          string
	OptionFunction      func(obj interface{}) error
	ServiceInstanceName string
}

func (o *BaseOption) GetServiceInstanceName() string {
	return o.ServiceInstanceName
}

func (o *BaseOption) GetServiceName() ServiceName {
	return o.ServiceName
}

func (o *BaseOption) GetOptionName() string {
	return o.OptionName
}

func (o *BaseOption) ApplyOptionFunction(obj interface{}) error {
	return o.OptionFunction(obj)
}
