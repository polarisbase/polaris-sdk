package model

type User struct {
	// ID is the unique identifier for the user
	ID string `json:"id"`
	// Username is the username for the user
	Username string `json:"username"`
	// BasicPasswordHash is the password for the user
	BasicPasswordHash string `json:"password"`
	// Email is the email for the user
	Email string `json:"email"`
	// FirstName is the first name for the user=
	FirstName string `json:"first_name"`
	// LastName is the last name for the user
	LastName string `json:"last_name"`
}
