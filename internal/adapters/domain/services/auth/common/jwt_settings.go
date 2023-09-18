package common

type JwtSettings struct {
	OnGetJwtSigningKeyByKid func(kid string) ([]byte, error)
	OnGetJwtSigningKey      func(id string, audience string, m map[string]string) ([]byte, error)
}

func (s *JwtSettings) GetJwtSigningKeyByKid(kid string) ([]byte, error) {

	if s.OnGetJwtSigningKeyByKid != nil {
		return s.OnGetJwtSigningKeyByKid(kid)
	}

	defaultKey := []byte("SecretYouShouldHide")

	return defaultKey, nil

}

func (s *JwtSettings) GetJwtSigningKey(id string, audience string, m map[string]string) ([]byte, error) {

	if s.OnGetJwtSigningKey != nil {
		return s.OnGetJwtSigningKey(id, audience, m)
	}

	defaultKey := []byte("SecretYouShouldHide")

	return defaultKey, nil
}
