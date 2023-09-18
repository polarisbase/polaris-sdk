package common

type Session interface {
	GetID() string
	GetUserID() string
	AsToken(settings *JwtSettings) (jwtTokenString string, err error)
	GetExpiresAt() int64
	GetNotBefore() int64
}
