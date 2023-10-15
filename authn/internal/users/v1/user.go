package v1

import (
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/shared"
)

func (a *Actions) FindUser(ticket shared.Ticket, request contracts.UserGetRequest) (response contracts.UserGetResponse, err error) {
	// Create the response
	response = contracts.UserGetResponse{}

	// Get the users
	err, user, ok := a.dep.UserStore.FindUserByID(ticket.GetContext(), request.ID)
	if !ok {
		return response, err
	}

	// Set the user
	response.User = contracts.UserDto{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	// Return the response
	return response, nil
}
