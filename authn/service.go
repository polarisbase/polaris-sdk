package authn

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/api"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/lib/persistence"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/lib/persistence/sqllite"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users"
)

type Service struct {
	// persist is the persistence layer
	persist persistence.SqlLite

	// fiberRouter is the router for the authentication API
	fiberRouter fiber.Router

	// authApiPrefix is the prefix for all authentication routes
	authApiPrefix string

	// userActionsProvider is the provider for user actions
	userActionsProvider *users.ActionsProvider

	// api is the authentication API
	api *api.Api
}

func New(fiberRouter fiber.Router, authApiPrefix string) *Service {
	// Create the authentication service
	s := &Service{
		authApiPrefix: authApiPrefix,
	}
	// Create the persistence layer
	s.persist = sqllite.New()
	// Connect to the persistence layer
	s.persist.Connect()
	// Create the user actions provider
	s.userActionsProvider = users.NewActionsProvider(s.persist)
	// Create a sub-router for the authentication API
	s.fiberRouter = fiberRouter.Group(s.authApiPrefix)
	// Create the authentication API
	s.api = api.New(s.fiberRouter, s.userActionsProvider)
	// Return the authentication service
	return s
}
