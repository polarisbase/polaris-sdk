package v1

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/model"
	"github.com/polarisbase/polaris-sdk/v2/shared"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (a *Actions) Login(ticket shared.Ticket, request contracts.LoginRequest) (responseOut contracts.LoginResponse, err error, ok bool) {
	// Create the response
	response := contracts.LoginResponse{}
	// Validate the request
	errValidateRequest, user, ok := a.validateLoginRequest(ticket, request)
	// If the request is not valid return the error
	if !ok {
		return response, errValidateRequest, false
	}
	// Create the ticket
	newTicket := shared.NewTicket(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		time.Until(time.Now().Add(time.Hour*12)),
		false,
	)
	// Set the response
	response.User = contracts.UserDto{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	response.SetTicket(newTicket)
	// Return the response
	return response, nil, true
}

func (a *Actions) validateLoginRequest(ticket shared.Ticket, request contracts.LoginRequest) (err error, user model.User, ok bool) {
	// Validate the email
	errValEmailResponse, userValEmailResponse, okValEmailResponse := a.validateLoginRequestEmail(ticket, request.Email)
	if !okValEmailResponse {
		_ = errValEmailResponse
		return err, user, false
	}
	// Validate the password
	errValPasswordResponse, okValPasswordResponse := a.validateLoginRequestPassword(ticket, request.Password, userValEmailResponse)
	if !okValPasswordResponse {
		_ = errValPasswordResponse
		return err, user, false
	}
	// Return nil
	return nil, userValEmailResponse, true
}

func (a *Actions) validateLoginRequestEmail(ticket shared.Ticket, email string) (err error, user model.User, ok bool) {
	errRes, userRes, okRes := a.dep.UserStore.FindUserByEmail(ticket.GetContext(), email)
	if !okRes {
		return errRes, userRes, false
	}

	return nil, userRes, true
}

func (a *Actions) validateLoginRequestPassword(ticket shared.Ticket, passwordIn string, user model.User) (err error, ok bool) {
	passwordInBytes := []byte(passwordIn)
	passwordHashBytes := []byte(user.BasicPasswordHash)

	// Comparing the password with the hash
	errComparingPassword := bcrypt.CompareHashAndPassword(passwordHashBytes, passwordInBytes)
	if errComparingPassword != nil {
		return fmt.Errorf("invalid password"), false
	} else {
		return nil, true
	}
}
