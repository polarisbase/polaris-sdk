package v1

import (
	"context"
	"fmt"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/lib/action"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/contracts"
	"golang.org/x/crypto/bcrypt"
)

func (a *Actions) Create(ticket action.Ticket, request contracts.CreateRequest) (contracts.CreateResponse, error) {
	// Create the response
	response := contracts.CreateResponse{}
	// Validate the request
	err := a.validateCreateRequest(request)
	if err != nil {
		return response, err
	}
	// hash the password
	// Hash the password
	hashIt := func(passwordIn string) string {
		password := []byte(passwordIn)

		// Hashing the password with the default cost of 10
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		return string(hashedPassword)

		//// Comparing the password with the hash
		//err = bcrypt.CompareHashAndPassword(hashedPassword, password)
		//fmt.Println(err) // nil means it is a match
	}
	hashedPassword := hashIt(request.Password)
	// Create the user
	if err, user, ok := a.dep.UserStore.CreateUser(context.Background(), request.Email, hashedPassword, map[string]string{
		"first_name": request.FirstName,
		"last_name":  request.LastName,
	}); ok {
		response.ID = user.ID
		response.Username = user.Username
		return response, nil
	} else {
		return response, err
	}

}

func (a *Actions) validateCreateRequest(request contracts.CreateRequest) error {
	// Validate the email
	err := a.validateCreateRequestEmail(request.Email)
	if err != nil {
		return err
	}
	// Validate the password
	err = a.validateCreateRequestPassword(request.Password)
	if err != nil {
		return err
	}

	// Return nil
	return nil
}

func (a *Actions) validateCreateRequestEmail(email string) error {
	// Validate the email
	if _, ok := a.dep.UserStore.CheckIfEmailIsAlreadyInUse(context.Background(), email); ok {
		return nil
	} else {
		return fmt.Errorf("Email is already in use")
	}
}

func (a *Actions) validateCreateRequestPassword(password string) error {
	return nil
}
