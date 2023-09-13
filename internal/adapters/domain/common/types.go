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
