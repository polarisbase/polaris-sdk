package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/shared"
)

func (a *Api) UserInfo(c *fiber.Ctx) error {

	// Ticket
	ticket := shared.GetTicket(c)
	if !ticket.IsAuthenticated() {
		c.Status(401)
		return c.JSON(shared.ApiResponse[contracts.UserGetResponse]{ // Change this to the correct response
			M: shared.ApiMetadata{
				StatusCode: 401,                            // Change this to the correct status code
				Status:     contracts.StatusUserGetFailure, // Change this to the correct status
			},
		})
	}

	// Get the user
	response, err := a.dep.UserActionsProvider.V1.FindUser(ticket, contracts.UserGetRequest{
		ID: ticket.GetUserID(),
	})

	// Check for errors
	if err != nil {
		fmt.Printf("Error: %s", err)
		c.Status(500)
		return c.JSON(shared.ApiResponse[contracts.UserGetResponse]{ // Change this to the correct response
			M: shared.ApiMetadata{
				Status: contracts.StatusUserGetFailure, // Change this to the correct status
			},
		})
	}

	// Create the response
	apiResponse := shared.ApiResponse[contracts.UserGetResponse]{ // Change this to the correct response
		M: shared.ApiMetadata{
			StatusCode: 200,                            // Change this to the correct status code
			Status:     contracts.StatusUserGetSuccess, // Change this to the correct status
		},
		Data: response,
	}

	fmt.Printf("Responce: %s", apiResponse)

	// Return the response
	return c.JSON(apiResponse)
}
