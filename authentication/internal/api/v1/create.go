package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users/contracts"
)

func (a *Api) Create(c *fiber.Ctx) error {

	// Ticket
	ticket := action.Ticket{}

	// Create the user
	response, err := a.dep.UserActionsProvider.V1.Create(ticket, contracts.CreateRequest{
		Username: "MichaelScott",
		Password: "Password",
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("Responce: %s", response)

	// Return nil
	return nil

}
