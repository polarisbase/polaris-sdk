package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/shared"
	"net/http"
)

func (a *Api) Login(c *fiber.Ctx) error {

	// Ticket
	ticket := action.Ticket{}

	// Get the request body from the context
	request := contracts.LoginRequest{}

	// Parse the request body
	err := c.BodyParser(&request)
	if err != nil {
		fmt.Printf("Error: %s", err)
		c.Status(http.StatusUnauthorized)
		return c.JSON(shared.ApiResponse[contracts.LoginResponse]{
			M: shared.ApiMetadata{
				Status: contracts.StatusLoginFailure,
			},
		})
	}

	// Login the user
	response, errLoginResponse, ok := a.dep.UserActionsProvider.V1.Login(ticket, request)
	if !ok {
		fmt.Printf("Error: %s", errLoginResponse)
		c.Status(http.StatusUnauthorized)
		return c.JSON(shared.ApiResponse[contracts.LoginResponse]{
			M: shared.ApiMetadata{
				Status: contracts.StatusLoginFailure,
			},
		})
	}

	apiResponse := shared.ApiResponse[contracts.LoginResponse]{
		M: shared.ApiMetadata{
			Status: contracts.StatusLoginSuccess,
		},
		Data: response,
	}

	fmt.Printf("Responce: %s", apiResponse)
	return c.JSON(apiResponse)
}
