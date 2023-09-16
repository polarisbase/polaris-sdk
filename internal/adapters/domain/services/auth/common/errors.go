package common

import "fmt"

var (
	ErrUserAlreadyExists          error = fmt.Errorf("user already exists")
	ErrUserNotFound               error = fmt.Errorf("user not found")
	ErrInvalidPassword            error = fmt.Errorf("invalid password")
	ErrInvalidEmail               error = fmt.Errorf("invalid email")
	ErrInvalidUsername            error = fmt.Errorf("invalid username")
	ErrInvalidSession             error = fmt.Errorf("invalid session")
	ErrInvalidCredentials         error = fmt.Errorf("invalid credentials")
	ErrInvalidSessionToken        error = fmt.Errorf("invalid session token")
	ErrInvalidSessionID           error = fmt.Errorf("invalid session id")
	ErrInvalidSessionRefreshToken error = fmt.Errorf("invalid session refresh token")
	ErrInvalidSessionProvider     error = fmt.Errorf("invalid session provider")
	ErrInvalidSessionUserID       error = fmt.Errorf("invalid session user id")
)
