package contracts

type CreateResponse struct {
	// ID is the unique identifier for the user
	ID string `json:"id"`
	// Username is the username for the user
	Username string `json:"username"`
	Redirect string `json:"redirect"`
}
