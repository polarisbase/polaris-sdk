package authweb

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/authweb/site"
)

type Service struct {
	*services.ServiceBase[common.OptionServiceAuthWebHtml]
	api  *api.Service
	auth *auth.Service
	site *site.Site
}

// ApplyLocalDefaults applies the default options.
func (s *Service) ApplyLocalDefaults() error {
	return nil
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
	s.ServiceBase = services.NewBase[common.OptionServiceAuthWebHtml](common.AUTHWEB_SERVICE, applicationName, s)

	// Apply service options.
	s.ServiceBase.ApplyOptions(options...)

	// Initialize api service
	if s.api == nil {
		s.api = api.New(applicationName, options...)
	}

	// Initialize auth service
	if s.auth == nil {
		s.auth = auth.New(
			applicationName,
			options...,
		)
	}

	// Initialize site
	s.site = site.New()

	// Add site
	s.api.DirectAccessFiberRouter().Get("/", s.site.EntryPoint())

	// Set the startup function.
	s.SetServiceStartupEntrypoint(func() error {
		_ = s.GetServiceInternalWaitGroup()

		return nil
	})

	// Return the service.
	return s
}
