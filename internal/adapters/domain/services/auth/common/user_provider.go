package common

type UserProvider interface {
	FindByEmail(email string) (User, error)
	FindByID(id string) (User, error)
	FindByUsername(username string) (User, error)
	NewBasic(email string, password string) (User, error)
}
