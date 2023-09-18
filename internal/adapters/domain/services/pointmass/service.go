package pointmass

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services"
	"sync"
)

type Service struct {
	*sync.RWMutex
	*services.ServiceBase[common.OptionServicePointmass]
	databases map[string]common.Database
}

// ApplyLocalDefaults applies the default options.
func (s *Service) ApplyLocalDefaults() error {

	// Initialize databases
	s.databases = make(map[string]common.Database)

	return nil
}

// NewDatabase creates a new database.
func (s *Service) NewDatabase(name string, driver common.DatabaseDriver) common.Database {

	// Initialize database
	db := driver.NewDatabase(name)

	// Add database to map
	s.databases[name] = db

	return db
}

// GetDatabase gets a database.
func (s *Service) GetDatabase(name string) common.Database {
	s.Lock()
	defer s.Unlock()
	return s.databases[name]
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
	s.ServiceBase = services.NewBase[common.OptionServicePointmass](common.POINTMASS_SERVICE, applicationName, s)

	// Apply service options.
	s.ServiceBase.ApplyOptions(options...)

	// Set the startup function.
	s.SetServiceStartupEntrypoint(func() error {
		_ = s.GetServiceInternalWaitGroup()

		return nil
	})

	// Return the service.
	return s
}
