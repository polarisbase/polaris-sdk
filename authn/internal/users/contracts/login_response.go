package contracts

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/v2/shared"
)

type LoginResponse struct {
	Redirect string  `json:"redirect,omitempty"`
	User     UserDto `json:"user,omitempty"`
	_ticket  shared.Ticket
}

func (r *LoginResponse) SetTicket(ticket shared.Ticket) {
	fmt.Printf("Set Ticket: %s", ticket.Print())
	r._ticket = ticket
}

func (r *LoginResponse) GetTicket() shared.Ticket {
	fmt.Printf("Get Ticket: %s", r._ticket.Print())
	return r._ticket
}
