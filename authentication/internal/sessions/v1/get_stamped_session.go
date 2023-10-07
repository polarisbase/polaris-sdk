package v1

import (
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/sessions/contracts"
)

func (a *Actions) GetStampedSession(ticket action.Ticket, request contracts.GetStampedSessionRequest) (responseOut contracts.GetStampedSessionResponse, errOut error, ok bool) {
	// Create the response
	response := contracts.GetStampedSessionResponse{}
	// Create the JWT
	
}
