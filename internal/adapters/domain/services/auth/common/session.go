package common

type Session interface {
	GetID() string
	GetUserID() string
	GetToken() string
	GetRefreshToken() string
	GetExpiresAt() int64
	GetNotBefore() int64
	GetCreatedAt() int64
	GetUpdatedAt() int64
	GetDeletedAt() int64

	ValidateToken(token string) bool
	ValidateRefreshToken(refreshToken string) bool
}
