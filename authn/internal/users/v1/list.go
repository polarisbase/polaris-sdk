package v1

import (
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/shared"
)

func (a *Actions) List(ticket shared.Ticket, request contracts.ListRequest) (response contracts.ListResponse, err error) {
	// Create the response
	response = contracts.ListResponse{}

	// Get the users
	err, users, ok := a.dep.UserStore.GetUsers(ticket.GetContext(), request.Limit, request.Offset)
	if !ok {
		return response, err
	}

	// Set the users
	for _, user := range users {
		response.Users = append(response.Users, contracts.UserDto{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
	}

	// Return the response
	return response, nil
}
