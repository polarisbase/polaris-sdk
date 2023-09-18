package common

type SessionProvider interface {
	NewSession(user User) (Session, error)
	NewSessionFromToken(token string) (Session, error)
	NewSessionFromRefreshToken(refreshToken string) (Session, error)
	FindByID(id string) (Session, error)
}