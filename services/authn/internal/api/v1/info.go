package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v3/shared"
)

func (a *Api) Info(c *fiber.Ctx) error {

	// Get ticket
	ticket := shared.GetTicket(c)

	// Do if user is authenticated
	return ticket.DoIfAuthenticated(func() error {
		c.Status(200)
		return c.JSON(map[string]string{
			"version": "v1",
		})
	}, func() error {
		c.Status(401)
		return c.JSON(map[string]string{})
	})

}
