package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/shared"
)

func (a *Api) Create(c *fiber.Ctx) error {

	// Ticket
	ticket := shared.GetTicket(c)

	// Get the request body from the context
	request := contracts.CreateRequest{}

	// Parse the request body
	err := c.BodyParser(&request)
	if err != nil {
		fmt.Printf("Error: %s", err)
		c.Status(500)
		return c.JSON(shared.ApiResponse[contracts.CreateResponse]{
			M: shared.ApiMetadata{
				Status: contracts.StatusCreateFailure,
			},
		})
	}

	// Create the user
	response, err := a.dep.UserActionsProvider.V1.Create(ticket, contracts.CreateRequest{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
		c.Status(500)
		return c.JSON(shared.ApiResponse[contracts.CreateResponse]{
			M: shared.ApiMetadata{
				Status: contracts.StatusCreateFailure,
			},
		})
	}

	apiResponse := shared.ApiResponse[contracts.CreateResponse]{
		M: shared.ApiMetadata{
			Status: contracts.StatusCreateSuccess,
		},
		Data: response,
	}

	fmt.Printf("Responce: %s", apiResponse)

	// set the client redirect (cr)
	if v := c.Query("cr"); v != "" {
		apiResponse.Data.Redirect = v
	}

	// Log response to console
	fmt.Printf("Responce: %s", apiResponse)

	// redirect
	// get the query params
	if c.Query("ur") == "true" {

		// get client
		client := c.Query("c")

		// lookup the redirect
		if client == "" {
			return c.Redirect("/")
		}

	}

	return c.JSON(apiResponse)

}
