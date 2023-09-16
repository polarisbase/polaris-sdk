package user

type BaseUser struct {
	ID             string
	Email          string
	Username       string
	PasswordHashed string
}

func (b BaseUser) GetID() string {
	return b.ID
}

func (b BaseUser) GetEmail() string {
	return b.Email
}

func (b BaseUser) GetUsername() string {
	return b.Username
}

func (b BaseUser) ValidatePassword(password string) bool {
	return password == b.PasswordHashed
}
