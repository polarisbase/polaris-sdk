package contracts

type CreateRequest struct {
	// FirstName is the first name of the user
	FirstName string `json:"first_name"`
	// LastName is the last name of the user
	LastName string `json:"last_name"`
	// Email is the username for the user
	Email string `json:"email"`
	// Password is the password for the user
	Password string `json:"password"`
}
