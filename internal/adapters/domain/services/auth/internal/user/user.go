package user

import "golang.org/x/crypto/bcrypt"

type BaseUser struct {
	ID             string
	Email          string
	Username       string
	PasswordHashed string
}

func (b BaseUser) GetID() string {
	return b.ID
}

func (b *BaseUser) SetEmail(email string) {
	b.Email = email
}

func (b BaseUser) GetEmail() string {
	return b.Email
}

func (b *BaseUser) SetUsername(username string) {
	b.Username = username
}

func (b BaseUser) GetUsername() string {
	return b.Username
}

func (b *BaseUser) SetPassword(password string) {

	// Convert the password string to a byte slice
	passwordAsBytes := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordAsBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	b.PasswordHashed = string(hashedPassword)

}

func (b BaseUser) ValidatePassword(password string) bool {

	// Convert the password string to a byte slice
	passwordAsBytes := []byte(password)

	// Convert the hashed password string to a byte slice
	hashedPasswordAsBytes := []byte(b.PasswordHashed)

	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPasswordAsBytes, passwordAsBytes)
	if err != nil {
		return false
	}

	return true
}
