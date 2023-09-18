package webhtml

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
	webCommon "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webhtml/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/webhtml/example"
)

type Service struct {
	*services.ServiceBase[common.OptionServiceWebHtml]
	apiService *api.Service
	site       webCommon.Site
}

// ApplyLocalDefaults applies the default options.
func (s *Service) ApplyLocalDefaults() error {
	return nil
}

func (s *Service) AddSite(site webCommon.Site) error {

	if s.site != nil {
		return fmt.Errorf("site already set")
	}

	s.site = site

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
	s.ServiceBase = services.NewBase[common.OptionServiceWebHtml](common.WEBHTML_SERVICE, applicationName, s)

	// Apply service options.
	s.ServiceBase.ApplyOptions(options...)

	// Initialize api service
	if s.apiService == nil {
		s.apiService = api.New(applicationName, options...)
	}

	// Set the startup function.
	s.SetServiceStartupEntrypoint(func() error {
		_ = s.GetServiceInternalWaitGroup()

		return nil
	})

	// Register example api
	example.NewExample(s.apiService)

	// Return the service.
	return s
}
