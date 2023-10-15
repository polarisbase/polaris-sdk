package fiber_middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/polarisbase/polaris-sdk/v3/shared"
)

func Tickets() func(c *fiber.Ctx) error {
	middleware := func(c *fiber.Ctx) error {

		// Set a custom header on all responses:
		c.Set("X-Tickets-Info", "Tickets-V1")

		// Set ticket in context
		c.Locals("ticket", &shared.TicketBase{
			UserId: "123",
		})

		// Go to next middleware:
		return c.Next()

	}
	return middleware
}
