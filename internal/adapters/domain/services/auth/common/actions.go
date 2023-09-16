package common

type Actions interface {
	RegisterNewUser(
		email string,
		password string,
	) (userID string, err error)

	SignIn(
		provider string,
		email string,
		password string,
	) (sessionID string, sessionToken string, err error)

	SignOut(
		sessionID string,
	) (err error)
}
