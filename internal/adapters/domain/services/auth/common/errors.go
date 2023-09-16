package common

import "fmt"

var PossibleErrors = Errors{
	UserAlreadyExists:          NewError("user already exists"),
	UserNotFound:               NewError("user not found"),
	InvalidPassword:            NewError("invalid password"),
	InvalidEmail:               NewError("invalid email"),
	InvalidUsername:            NewError("invalid username"),
	InvalidSession:             NewError("invalid session"),
	InvalidCredentials:         NewError("invalid credentials"),
	InvalidSessionToken:        NewError("invalid session token"),
	InvalidSessionID:           NewError("invalid session id"),
	InvalidSessionRefreshToken: NewError("invalid session refresh token"),
	InvalidSessionProvider:     NewError("invalid session provider"),
	InvalidSessionUserID:       NewError("invalid session user id"),
}

func NewError(s string) error {
	return fmt.Errorf(s)
}

type Errors struct {
	UserAlreadyExists          error
	UserNotFound               error
	InvalidPassword            error
	InvalidEmail               error
	InvalidUsername            error
	InvalidSession             error
	InvalidCredentials         error
	InvalidSessionToken        error
	InvalidSessionID           error
	InvalidSessionRefreshToken error
	InvalidSessionProvider     error
	InvalidSessionUserID       error
}
