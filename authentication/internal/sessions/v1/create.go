package v1

import (
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/sessions/contracts"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/sessions/model"
)

func (a *Actions) Create(ticket action.Ticket, request contracts.CreateRequest) (responseOut contracts.CreateResponse, errOut error, ok bool) {
	// Create the response
	response := contracts.CreateResponse{}
	// Validate the request
	if err, ok := a.validateCreateRequest(request); !ok {
		return response, err, false
	}
	// Create the session
	newSession := model.Session{
		UserID: request.UserID,
	}
	// Create the session
	if err, session, ok := a.dep.SessionStore.CreateSession(ticket.GetTicketContext(), newSession); !ok {
		return response, err, false
	} else {
		response.ID = session.ID
		return response, err, true
	}
}

func (a *Actions) validateCreateRequest(request contracts.CreateRequest) (err error, ok bool) {
	return nil, true
}
