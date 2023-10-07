package v1

import (
	"fmt"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users/contracts"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users/model"
	"golang.org/x/crypto/bcrypt"
)

func (a *Actions) Login(ticket action.Ticket, request contracts.LoginRequest) (responseOut contracts.LoginResponse, err error, ok bool) {
	// Create the response
	response := contracts.LoginResponse{}
	// Validate the request
	errValidateRequest, user, ok := a.validateLoginRequest(ticket, request)
	if !ok {
		return response, errValidateRequest, false
	}
	// Login the user
	_ = user
	// Create the token
	_ = user
	// Set the response
	response.AccessCode = "access_code"
	response.ExpiresIn = 3600
	// Return the response
	return response, nil, true
}

func (a *Actions) validateLoginRequest(ticket action.Ticket, request contracts.LoginRequest) (err error, user model.User, ok bool) {
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
	return nil, user, true
}

func (a *Actions) validateLoginRequestEmail(ticket action.Ticket, email string) (err error, user model.User, ok bool) {
	errRes, userRes, okRes := a.dep.UserStore.FindUserByEmail(ticket.GetTicketContext(), email)
	if !okRes {
		return errRes, userRes, false
	}

	return nil, userRes, true
}

func (a *Actions) validateLoginRequestPassword(ticket action.Ticket, passwordIn string, user model.User) (err error, ok bool) {
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
