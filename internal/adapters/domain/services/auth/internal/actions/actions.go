package actions

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/internal/session"
)

type Actions struct {
	users       common.UserProvider
	sessions    common.SessionProvider
	jwtSettings *common.JwtSettings
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
		return "", "", common.PossibleErrors.InvalidCredentials
	}

	session, err := a.sessions.NewSession(user)
	if err != nil {
		return "", "", err
	}

	token, err := session.AsToken(a.jwtSettings)
	if err != nil {
		return "", "", err
	}

	return session.GetID(), token, nil
}

func (a *Actions) SignOut(sessionID string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *Actions) ValidateSession(sessionToken string) (sessionOut common.Session, err error) {
	sessionOut, valid, err := session.Verify(a.jwtSettings, sessionToken)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, common.PossibleErrors.InvalidSession
	}

	return sessionOut, nil
}

func New(
	users common.UserProvider,
	sessions common.SessionProvider,
	jwtSettings *common.JwtSettings,
) *Actions {
	a := &Actions{}

	a.users = users
	a.sessions = sessions
	a.jwtSettings = jwtSettings

	return a
}
