package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services"
)

type Service struct {
	*services.ServiceBase[common.OptionServiceApi]
	fiber              *fiber.App
	fiberRouterPrefix  string
	fiberRouter        fiber.Router
	fiberPortsToListen []string
}

// ApplyLocalDefaults applies the default options.
func (s *Service) ApplyLocalDefaults() error {
	s.fiberPortsToListen = []string{":3000"}
	s.fiberRouterPrefix = ""
	return nil
}

// SetFiberPortsToListen sets the fiber ports to listen.
func (s *Service) SetFiberPortsToListen(portsToListen []string) {
	s.fiberPortsToListen = portsToListen
}

// SetFiberPathPrefix sets the fiber path prefix.
func (s *Service) SetFiberRouterPrefix(prefix string) {
	s.fiberRouterPrefix = prefix
}

// DirectAccessFiber returns the fiber router.
func (s *Service) DirectAccessFiber() *fiber.App {
	return s.fiber
}

// DirectAccessFiberRouter returns the fiber router.
func (s *Service) DirectAccessFiberRouter() fiber.Router {
	return s.fiberRouter
}

// SetFiberApp sets the fiber app.
func (s *Service) SetFiberApp(app *fiber.App) {
	s.fiber = app
}

// SetAsNonExacutable sets the service as non executable.
func (s *Service) SetAsNonExacutable(isNonExecutable bool) {
	s.ServiceBase.SetAsNonExacutable(isNonExecutable)
}

// Start starts the service.
func (s *Service) Start() error {
	fmt.Printf("STARTING SERVICE: %s\n", s.GetName())
	return s.ServiceBase.Start()
}

// New creates a new Service.
func New(applicationName string, options ...common.Option) *Service {

	// Create the service.
	s := &Service{}

	// Apply local the defaults.
	if err := s.ApplyLocalDefaults(); err != nil {
		panic(err)
	}

	// Set the base service
	s.ServiceBase = services.NewBase[common.OptionServiceApi](common.API_FIBER_SERVICE, options...)

	// Create the fiber router.
	if s.fiber == nil {
		s.fiber = fiber.New()
	}

	// Set the fiber router prefix.
	if s.fiberRouterPrefix == "" || s.fiberRouterPrefix == "/" {
		s.fiberRouter = s.fiber
	} else {
		s.fiberRouter = s.fiber.Group(s.fiberRouterPrefix)
	}

	// Set the service name.
	s.SetName(applicationName)

	// Set the startup function.
	s.SetServiceStartupEntrypoint(func() error {

		internalWaitGroup := s.GetServiceInternalWaitGroup()

		for _, portToListen := range s.fiberPortsToListen {

			internalWaitGroup.Add(1)

			go func(portToListen string) {
				defer internalWaitGroup.Done()

				s.OnInternalError(s.fiber.Listen(portToListen))

			}(portToListen)

		}

		return nil

	})

	// Return the service.
	return s

}
