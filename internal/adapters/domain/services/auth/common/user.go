package common

type User interface {
	GetID() string
	GetEmail() string
	GetUsername() string
	ValidatePassword(password string) bool
}
