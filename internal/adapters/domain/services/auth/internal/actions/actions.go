package actions

import "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"

type Actions struct {
	users    common.UserProvider
	sessions common.SessionProvider
}

func (a *Actions) RegisterNewUser(email string, password string) (userID string, err error) {

	user, err := a.users.NewBasic(email, password)

	if err != nil {
		return "", err
	}

	return user.GetID(), nil
}

func (a *Actions) SignIn(provider string, email string, password string) (sessionID string, sessionToken string, err error) {
	user, err := a.users.FindByEmail(email)
	if err != nil {
		return "", "", err
	}

	if !user.ValidatePassword(password) {
		return "", "", common.ErrInvalidCredentials
	}

	session, err := a.sessions.NewSession(user)
	if err != nil {
		return "", "", err
	}

	return session.GetID(), session.GetToken(), nil
}

func (a *Actions) SignOut(sessionID string) (err error) {
	//TODO implement me
	panic("implement me")
}

func New(
	users common.UserProvider,
	sessions common.SessionProvider,
) *Actions {
	a := &Actions{}

	a.users = users
	a.sessions = sessions

	return a
}
