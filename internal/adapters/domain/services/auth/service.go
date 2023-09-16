package auth

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
	userCommon "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/internal/actions"
	authApi "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/internal/api"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/internal/session"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/internal/user"
)

type Service struct {
	*services.ServiceBase[common.OptionServiceApi]
	apiService      *api.Service
	api             *authApi.Api
	userProvider    userCommon.UserProvider
	sessionProvider userCommon.SessionProvider
	actions         userCommon.Actions
}

func (s *Service) GetActions() userCommon.Actions {
	return s.actions
}

// ApplyLocalDefaults applies the default options.
func (s *Service) ApplyLocalDefaults() error {

	// Initialize user provider
	s.userProvider = user.NewProvider()

	// Initialize session provider
	s.sessionProvider = session.NewProvider()

	return nil
}

// Start starts the service.
func (s *Service) Start() error {
	fmt.Printf("STARTING SERVICE: %s\n", s.GetServiceInstanceName())
	return s.ServiceBase.Start()
}

// New creates a new service.
func New(applicationName string, options ...common.Option) *Service {

	// Create the service.
	s := &Service{}

	// Apply local defaults.
	if err := s.ApplyLocalDefaults(); err != nil {
		panic(err)
	}

	// Set the base service.
	s.ServiceBase = services.NewBase[common.OptionServiceApi](common.AUTH_SERVICE, applicationName, s)

	// Apply service options.
	s.ServiceBase.ApplyOptions(options...)

	// Initialize actions
	s.actions = actions.New(s.userProvider, s.sessionProvider)

	// Initialize api service
	if s.apiService == nil {
		s.apiService = api.New(applicationName, options...)
	}

	// Initialize api
	if s.api == nil {
		s.api = authApi.New(s.apiService, s.actions)
	}

	// Set the startup function.
	s.SetServiceStartupEntrypoint(func() error {
		_ = s.GetServiceInternalWaitGroup()

		return nil
	})

	// Return the service.
	return s
}
