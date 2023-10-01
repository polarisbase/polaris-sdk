package v1

import (
	"context"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users/contracts"
)

func (a *Actions) Create(ticket action.Ticket, request contracts.CreateRequest) (contracts.CreateResponse, error) {
	// Create the response
	response := contracts.CreateResponse{}
	// Validate the request
	err := a.validateCreateRequest(request)
	if err != nil {
		return response, err
	}
	// Create the user
	if err, user, ok := a.dep.UserStore.CreateUser(context.Background(), request.Username, request.Password); ok {
		response.ID = user.ID
		response.Username = user.Username
		return response, nil
	} else {
		return response, err
	}

}

func (a *Actions) validateCreateRequest(request contracts.CreateRequest) error {
	// Validate the email
	err := a.validateEmail(request.Username)
	if err != nil {
		return err
	}
	// Validate the password
	err = a.validatePassword(request.Password)
	if err != nil {
		return err
	}

	// Return nil
	return nil
}

func (a *Actions) validateEmail(email string) error {
	// Validate the email
	if err, ok := a.dep.UserStore.CheckIfEmailIsAlreadyInUse(context.Background(), email); ok {
		return nil
	} else {
		return err
	}
}

func (a *Actions) validatePassword(password string) error {
	return nil
}
