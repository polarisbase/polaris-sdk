package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
)

type ServiceBase[serviceOptionType common.OptionService] struct {
	serviceSysName              common.ServiceName
	serviceImplementation       interface{}
	serviceName                 string
	serviceEventChannel         chan interface{}
	onInternalApplyOptionsError func(err error) (shouldContinue bool, errOut error)
	useHeartbeat                bool
	onInternalError             func(err error)
	serviceInternalWaitGroup    *sync.WaitGroup
	serviceStartupEntrypoint    func() error
	nonExacutable               bool
	allowRunAsNonExecutable     bool
}

// SetName sets the name of the service.
func (s *ServiceBase[serviceOptionType]) SetName(name string) {
	s.serviceName = name
}

// GetName gets the name of the service.
func (s *ServiceBase[serviceOptionType]) GetName() string {
	return s.serviceName
}

// GetServiceSystemName gets the system name of the service.
func (s *ServiceBase[serviceOptionType]) GetServiceSystemName() common.ServiceName {
	return s.serviceSysName
}

// getServiceImplementation gets the service implementation.
func (s *ServiceBase[serviceOptionType]) getServiceImplementation() interface{} {
	return s.serviceImplementation
}

// ApplyDefaults applies the default options to the service.
func (s *ServiceBase[serviceOptionType]) ApplyDefaults() error {

	s.useHeartbeat = true

	s.allowRunAsNonExecutable = false

	s.nonExacutable = false

	s.serviceEventChannel = make(chan interface{}, 100)

	s.onInternalError = func(err error) {
		fmt.Printf("DEFAULT ERROR HANDLER: %s\n", err.Error())
	}

	s.onInternalApplyOptionsError = func(err error) (shouldContinue bool, errOut error) {
		fmt.Printf("DEFAULT APPLY OPTIONS ERROR HANDLER: %s\n", err.Error())
		return false, err
	}

	s.serviceInternalWaitGroup = &sync.WaitGroup{}

	s.serviceStartupEntrypoint = func() error {
		fmt.Printf("DEFAULT STARTUP ENTRYPOINT\n")
		return nil
	}

	return nil
}

// SetServiceEventChannel sets the service event stream.
func (s *ServiceBase[serviceOptionType]) SetServiceEventChannel(serviceEventChannel chan interface{}) {
	s.serviceEventChannel = serviceEventChannel
}

// GetServiceEventChannel gets the service event stream.
func (s *ServiceBase[serviceOptionType]) GetServiceEventChannel() chan interface{} {
	return s.serviceEventChannel
}

// SetAsNonExacutable sets the service as non executable.
func (s *ServiceBase[serviceOptionType]) SetAsNonExacutable(isNonExecutable bool) {
	s.nonExacutable = isNonExecutable
}

// ApplyOptions applies the given options to the service.
func (s *ServiceBase[serviceOptionType]) ApplyOptions(options ...common.Option) error {
	for _, opt := range options {

		if opt == nil {
			continue
		}

		if opt.GetOptionName() == "" {
			continue
		}

		if opt.GetServiceName() != s.serviceSysName {
			continue
		}

		if optServiceOption, ok := opt.(common.OptionService); ok {
			if err := optServiceOption.ApplyOptionFunction(s); err != nil {
				return err
			}
		}

		if optServiceOption, ok := opt.(serviceOptionType); ok {
			if err := optServiceOption.ApplyOptionFunction(s.getServiceImplementation()); err != nil {
				return err
			}
		}

	}
	return nil
}

// UseHeartbeat returns true if the service should use a heartbeat.
func (s *ServiceBase[serviceOptionType]) UseHeartbeat(ans bool) {
	s.useHeartbeat = ans
}

// SetServiceInternalWaitGroup sets the service internal wait group.
func (s *ServiceBase[serviceOptionType]) SetServiceInternalWaitGroup(serviceInternalWaitGroup *sync.WaitGroup) {
	s.serviceInternalWaitGroup = serviceInternalWaitGroup
}

// GetServiceInternalWaitGroup gets the service internal wait group.
func (s *ServiceBase[serviceOptionType]) GetServiceInternalWaitGroup() *sync.WaitGroup {
	return s.serviceInternalWaitGroup
}

// SetOnInternalApplyOptionsError sets the OnInternalApplyOptionsError function.
func (s *ServiceBase[serviceOptionType]) SetOnInternalApplyOptionsError(onInternalApplyOptionsError func(err error) (shouldContinue bool, errOut error)) {
	s.onInternalApplyOptionsError = onInternalApplyOptionsError
}

// OnInternalApplyOptionsError calls the OnInternalApplyOptionsError function.
func (s *ServiceBase[serviceOptionType]) OnInternalApplyOptionsError(errIn error) (shouldContinue bool, err error) {
	if s.onInternalApplyOptionsError != nil {
		return s.onInternalApplyOptionsError(errIn)
	}
	return false, errIn
}

// SetOnInternalError sets the OnInternalError function.
func (s *ServiceBase[serviceOptionType]) SetOnInternalError(onInternalError func(err error)) {
	s.onInternalError = onInternalError
}

// OnInternalError calls the OnInternalError function.
func (s *ServiceBase[serviceOptionType]) OnInternalError(errIn error) (shouldContinue bool, err error) {
	if s.onInternalError != nil {
		s.onInternalError(errIn)
	}
	return false, errIn
}

// SetServiceStartupEntrypoint sets the service startup entrypoint.
func (s *ServiceBase[serviceOptionType]) SetServiceStartupEntrypoint(serviceStartupEntrypoint func() error) {
	s.serviceStartupEntrypoint = serviceStartupEntrypoint
}

// GetServiceStartupEntrypoint gets the service startup entrypoint.
func (s *ServiceBase[serviceOptionType]) GetServiceStartupEntrypoint() func() error {
	return s.serviceStartupEntrypoint
}

// Start starts the service. This is a blocking call.
func (s *ServiceBase[serviceOptionType]) Start() error {

	if s.serviceStartupEntrypoint == nil {
		return fmt.Errorf("service startup entrypoint is nil")
	}

	if s.nonExacutable {
		if !s.allowRunAsNonExecutable {
			return fmt.Errorf("service is marked as non executable")
		} else {
			fmt.Printf("WARNING: Service %s is marked as non executable (but is allowed to run)\n", s.serviceName)
		}
	}

	startErr := s.serviceStartupEntrypoint()
	if startErr != nil {
		return startErr
	}

	time.Sleep(1 * time.Second)

	s.GetServiceInternalWaitGroup().Wait()

	return nil
}

func NewBase[serviceOptionType common.OptionService](serviceSysName common.ServiceName, serviceImplementation interface{}) *ServiceBase[serviceOptionType] {
	s := &ServiceBase[serviceOptionType]{
		serviceSysName: serviceSysName,
		serviceImplementation: serviceImplementation,
	}

	// Apply the default options.
	if err := s.ApplyDefaults(); err != nil {
		shouldContinue, err := s.OnInternalApplyOptionsError(err)
		if !shouldContinue {
			fmtErr := fmt.Errorf("failed to apply defaults for internal api service %s: %w", s.serviceSysName, err)
			fmt.Println(fmtErr.Error())
			return nil
		}
	}

	return s
}
