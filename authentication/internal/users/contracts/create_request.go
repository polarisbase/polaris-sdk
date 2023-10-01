package contracts

type CreateRequest struct {
	// Username is the username for the user
	Username string `json:"username"`
	// Password is the password for the user
	Password string `json:"password"`
}
