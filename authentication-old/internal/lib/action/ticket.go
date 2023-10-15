package action

import "context"

type Ticket struct {
}

func (t *Ticket) GetTicketContext() context.Context {
	return context.Background()
}
