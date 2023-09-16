package session

import "time"

type BaseSession struct {
	ID           string
	UserID       string
	Token        string
	RefreshToken string
	ExpiresAt    time.Time
	NotBefore    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func (b BaseSession) GetID() string {
	return b.ID
}

func (b BaseSession) GetUserID() string {
	return b.UserID
}

func (b BaseSession) GetToken() string {
	return b.Token
}

func (b BaseSession) GetRefreshToken() string {
	return b.RefreshToken
}

func (b BaseSession) GetExpiresAt() int64 {
	return b.ExpiresAt.Unix()
}

func (b BaseSession) GetNotBefore() int64 {
	return b.NotBefore.Unix()
}

func (b BaseSession) GetCreatedAt() int64 {
	return b.CreatedAt.Unix()
}

func (b BaseSession) GetUpdatedAt() int64 {
	return b.UpdatedAt.Unix()
}

func (b BaseSession) GetDeletedAt() int64 {
	return b.DeletedAt.Unix()
}

func (b BaseSession) ValidateToken(token string) bool {
	//TODO implement me
	panic("implement me")
}

func (b BaseSession) ValidateRefreshToken(refreshToken string) bool {
	//TODO implement me
	panic("implement me")
}
