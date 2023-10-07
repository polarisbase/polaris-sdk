package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/shared"
)

func (a *Api) Create(c *fiber.Ctx) error {

	// Ticket
	ticket := action.Ticket{}

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
	return c.JSON(apiResponse)

}
