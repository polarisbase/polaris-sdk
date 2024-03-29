package v1

import (
	"github.com/polarisbase/polaris-sdk/v3/services/templates/template_1/internal/info/contracts"
	"github.com/polarisbase/polaris-sdk/v3/services/templates/template_1/internal/info/model"
	"github.com/polarisbase/polaris-sdk/v3/shared"
)

func (a *Actions) Create(ticket shared.Ticket, request contracts.CreateRequest) (contracts.CreateResponse, error) {
	// Create the response
	response := contracts.CreateResponse{}

	// Create in the store
	res, err, ok := a.dep.InfoStore.CreateInfo(ticket.GetContext(), model.Info{})
	if err != nil {
		return response, err
	}
	if !ok {
		return response, nil
	}

	// Set the response
	response.ID = res.ID

	// Return the response
	return response, nil
}
