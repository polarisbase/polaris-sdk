package common

type ServiceName string

type Option interface {
	GetServiceName() ServiceName
	GetOptionName() string
	ApplyOptionFunction(obj interface{}) error
}

type BaseOption struct {
	ServiceName    ServiceName
	OptionName     string
	OptionFunction func(obj interface{}) error
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
