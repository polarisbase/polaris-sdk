package ports

import (
	"sync"

	"github.com/polarisbase/polaris-go/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-go/internal/adapters/domain/services"
	"github.com/polarisbase/polaris-go/internal/adapters/domain/services/api"
)

type Service interface {
	GetServiceSystemName() common.ServiceName
	SetName(name string)
	GetName() string

	SetOnInternalApplyOptionsError(onInternalApplyOptionsError func(err error) (shouldContinue bool, errOut error))
	OnInternalApplyOptionsError(errIn error) (shouldContinue bool, err error)

	ApplyDefaults() error
	ApplyOptions(options ...common.Option) error

	SetServiceEventChannel(serviceEventChannel chan interface{})
	GetServiceEventChannel() chan interface{}

	UseHeartbeat(ans bool)

	SetOnInternalError(onInternalError func(err error))
	OnInternalError(errIn error) (shouldContinue bool, err error)

	SetServiceInternalWaitGroup(serviceInternalWaitGroup *sync.WaitGroup)
	GetServiceInternalWaitGroup() *sync.WaitGroup

	SetServiceStartupEntrypoint(serviceStartupEntrypoint func() error)
	GetServiceStartupEntrypoint() func() error

	Start() error
}

// Ensure that Service implements the Service interface.
var _ Service = (*services.ServiceBase[common.OptionService])(nil)
var _ Service = (*api.Service)(nil)
