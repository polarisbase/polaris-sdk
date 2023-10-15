package authn

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v3/lib/persist/document"
	"github.com/polarisbase/polaris-sdk/v3/services/authn/internal/api"
)

type Service struct {
	// fiberRouter is the router for the authentication API
	fiberRouter fiber.Router
	// authApiPrefix is the prefix for all authentication routes
	authApiPrefix string
	// documentStore is the document store
	documentStore document.Store
	// api is the authentication API
	api *api.Api
}

func New(fiberRouter fiber.Router, authApiPrefix string, documentStore document.Store) *Service {
	// Create the authentication service
	s := &Service{
		authApiPrefix: authApiPrefix,
	}
	// Create the document store
	s.documentStore = documentStore
	// Create a sub-router for the authentication API
	s.fiberRouter = fiberRouter.Group(s.authApiPrefix)
	// Create the authentication API
	s.api = api.New(s.fiberRouter)
	// Return the authentication service
	return s
}
