package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/shared"
)

func (a *Api) List(c *fiber.Ctx) error {
	// Ticket
	ticket := action.Ticket{}

	//// Get the request body from the context
	//request := contracts.CreateRequest{}
	//
	//// Parse the request body
	//err := c.BodyParser(&request)
	//if err != nil {
	//	fmt.Printf("Error: %s", err)
	//	c.Status(500)
	//	return c.JSON(shared.ApiResponse[contracts.CreateResponse]{
	//		M: shared.ApiMetadata{
	//			Status: contracts.StatusCreateFailure,
	//		},
	//	})
	//}

	// Get the users
	response, err := a.dep.UserActionsProvider.V1.List(c, ticket, contracts.ListRequest{})
	if err != nil {
		fmt.Printf("Error: %s", err)
		c.Status(500)
		return c.JSON(shared.ApiResponse[contracts.CreateResponse]{
			M: shared.ApiMetadata{
				Status: contracts.StatusListFailure,
			},
		})
	}

	// Create the response
	apiResponse := shared.ApiResponse[contracts.ListResponse]{
		M: shared.ApiMetadata{
			Status: contracts.StatusListSuccess,
		},
		Data: response,
	}

	// Return the response
	return c.JSON(apiResponse)
}
