package session

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"
	"time"
)

type BaseSession struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	KID       string    `json:"kid"`
	UserID    string    `json:"user_id"`
	Issuer    string    `json:"issuer"`
	Audience  string    `json:"audience"`
	ExpiresAt time.Time `json:"expires_at"`
	NotBefore time.Time `json:"not_before"`
	IssuedAt  time.Time `json:"issued_at"`
}

func (s *BaseSession) GetID() string {
	return s.ID
}

func (s *BaseSession) GetUserID() string {
	return s.UserID
}

func (s *BaseSession) GetExpiresAt() int64 {
	return s.ExpiresAt.Unix()
}

func (s *BaseSession) GetNotBefore() int64 {
	return s.NotBefore.Unix()
}

func (s *BaseSession) AsToken(settings *common.JwtSettings) (jwtTokenString string, err error) {

	// New creates a new JWT token with the given signing method.
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the "kid" (key ID) header claim.
	token.Header["kid"] = s.KID

	// Claims represents the standard JWT claims.
	claims := token.Claims.(jwt.MapClaims)

	// Set the "iss" (issuer) claim.
	claims["iss"] = s.Issuer

	// Set the "sub" (subject) claim.
	claims["sub"] = s.UserID

	// Set the "aud" (audience) claim.
	claims["aud"] = s.Audience

	// Set the "exp" (expiration time) claim.
	claims["exp"] = s.ExpiresAt.Unix()

	// Set the "nbf" (not before) claim.
	claims["nbf"] = s.NotBefore.Unix()

	// Set the "iat" (issued at) claim.
	claims["iat"] = s.IssuedAt.Unix()

	// Set the "jti" (JWT ID) claim.
	claims["jti"] = s.ID

	// Sign and get the complete encoded token as a string using the secret
	// signing key.

	var signingKey []byte

	if keyByKid, err := settings.GetJwtSigningKeyByKid(s.KID); err == nil {
		signingKey = keyByKid
	} else {

		if key, err := settings.GetJwtSigningKey(s.UserID, s.Audience, map[string]string{}); err == nil {
			signingKey = key
		} else {
			return "", err
		}

	}

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Verify(settings *common.JwtSettings, tokenString string) (session *BaseSession, valid bool, err error) {

	s := &BaseSession{}

	// Parse parses, validates, and returns a token. KeyFunc will receive the
	// parsed token and should return the key for validating. If everything is
	// kosher, it will return a nil error.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Check the signing method.
		if token.Method != jwt.SigningMethodEdDSA && token.Method != jwt.SigningMethodHS256 {
			return nil, common.PossibleErrors.InvalidSigningAlgorithm
		}

		// Check for the "kid" (key ID) claim.
		if kid, ok := token.Header["kid"]; !ok {
			return nil, common.PossibleErrors.MissingKeyID
		} else {
			s.KID = kid.(string)
		}

		// Check for the "iss" (issuer) claim.
		if iss, ok := token.Claims.(jwt.MapClaims)["iss"]; !ok {
			return nil, common.PossibleErrors.MissingIssuer
		} else {
			s.Issuer = iss.(string)
		}

		// Check for the "sub" (subject) claim.
		if sub, ok := token.Claims.(jwt.MapClaims)["sub"]; !ok {
			return nil, common.PossibleErrors.MissingSubject
		} else {
			s.UserID = sub.(string)
		}

		// Check for the "aud" (audience) claim.
		if aud, ok := token.Claims.(jwt.MapClaims)["aud"]; !ok {
			return nil, common.PossibleErrors.MissingAudience
		} else {
			s.Audience = aud.(string)
		}

		// Check for the "exp" (expiration time) claim.
		if exp, ok := token.Claims.(jwt.MapClaims)["exp"]; !ok {
			return nil, common.PossibleErrors.MissingExpirationTime
		} else {
			s.ExpiresAt = time.Unix(int64(exp.(float64)), 0)
		}

		// Check for the "nbf" (not before) claim.
		if nbf, ok := token.Claims.(jwt.MapClaims)["nbf"]; !ok {
			return nil, common.PossibleErrors.MissingNotBefore
		} else {
			s.NotBefore = time.Unix(int64(nbf.(float64)), 0)
		}

		// Check for the "iat" (issued at) claim.
		if iat, ok := token.Claims.(jwt.MapClaims)["iat"]; !ok {
			return nil, common.PossibleErrors.MissingIssuedAt
		} else {
			s.IssuedAt = time.Unix(int64(iat.(float64)), 0)
		}

		// Check for the "jti" (JWT ID) claim.
		if jti, ok := token.Claims.(jwt.MapClaims)["jti"]; !ok {
			return nil, common.PossibleErrors.MissingJWTID
		} else {
			s.ID = jti.(string)
		}

		// Get the signing key.
		var signingKey []byte

		if keyByKid, err := settings.GetJwtSigningKeyByKid(s.KID); err == nil {
			signingKey = keyByKid
		} else {

			if key, err := settings.GetJwtSigningKey(s.UserID, s.Audience, map[string]string{}); err == nil {
				signingKey = key
			} else {
				return nil, err
			}

		}

		// Return the signing key.
		return signingKey, nil
	})
	if err != nil {
		return nil, false, err
	}

	// Check if the token is valid.
	if !token.Valid {
		return nil, false, jwt.ErrSignatureInvalid
	}

	// Return the claims.
	return s, true, nil
}

func NewBasic(userId string) *BaseSession {
	return &BaseSession{
		ID:        uuid.New().String(),
		KID:       "-",
		Issuer:    "-",
		UserID:    userId,
		Audience:  "-",
		ExpiresAt: time.Now().Add(time.Hour * 6),
		NotBefore: time.Now(),
		IssuedAt:  time.Now(),
	}
}
