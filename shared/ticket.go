package shared

import "github.com/gofiber/fiber/v2"

type Ticket interface {
	GetUserId() string
	DoIfAuthenticated(authenticated func() error, unauthenticated func() error) error
}

type TicketBase struct {
	UserId string
}

func (t TicketBase) GetUserId() string {
	return t.UserId
}

func (t TicketBase) DoIfAuthenticated(authenticated func() error, unauthenticated func() error) error {
	if t.UserId != "" {
		return authenticated()
	} else {
		return unauthenticated()
	}
}

func GetTicket(c interface{}) (ticket Ticket) {

	// Check if the context is nil
	if c == nil {
		return nil
	}

	// Get the ticket from the context if of type Fiber context
	if v, ok := c.(*fiber.Ctx); ok {
		found := v.Locals("ticket")
		if found != nil {
			return found.(Ticket)
		}
	}

	// return the default ticket
	return nil

}
