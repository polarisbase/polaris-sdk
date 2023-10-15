package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/shared"
)

func (a *Api) Logout(c *fiber.Ctx) error {

	// Ticket
	ticket := shared.GetTicket(c)

	// Delete the ticket
	if err := ticket.DeleteTicketFromSession(c); err != nil {
		c.Status(500)
		return c.JSON(shared.ApiResponse[contracts.LogoutResponse]{
			M: shared.ApiMetadata{
				StatusCode: 500,
				Status:     contracts.StatusLogoutFailure,
			},
		})
	}

	// Create the response
	apiResponse := shared.ApiResponse[contracts.LogoutResponse]{
		M: shared.ApiMetadata{
			StatusCode: 200,
			Status:     contracts.StatusLogoutSuccess,
		},
	} // Change this to the correct response

	// Return the response
	return c.JSON(apiResponse)

}
