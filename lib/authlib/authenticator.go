package authlib

import (
	"sync"

	"context"
)

type basicAuthenticator struct {
	keyProvider string
	privateKey string
	publicKeys map[int]string
	sync.Mutex
}

func NewBasicAuthenticator(keyProvider KeyProvider) *basicAuthenticator {
	key, err := keyProvider.GetPrivateKey()
	if err != nil {
		panic(err)
	}
	return &basicAuthenticator{
		privateKey: key,
	}
}

func (a *basicAuthenticator) GenerateServiceToken(serviceID int) string {
	return ""
}

func (a *basicAuthenticator) GenerateServiceTokenWithUserData(serviceID int, ctx context.Context) string {
	return ""
}

func (a *basicAuthenticator) ParseUserToken(token string) (UserToken, error) {
	return UserToken{}, nil
}

func (a *basicAuthenticator) ParseApplicationToken(token string) (ApplicationToken, error) {
	return ApplicationToken{}, nil
}

/*
func DecodeToken(token string) error {
	var (
		rawHeader  []byte
		rawPayload []byte
	)
	tokenSplit := strings.Split(token, ".")
	if len(tokenSplit) != 3 {
		return errorlib.Errf("Token is invalid, expected 3 parts seperated by '.' but got %d", len(tokenSplit))
	}

	_, err := base64.RawURLEncoding.Decode(rawHeader, []byte(tokenSplit[0]))
	if err != nil {
		return err
	}
	var header JWTHeader
	err = json.Unmarshal(rawHeader, &header)
	if err != nil {
		return err
	}

	_, err = base64.RawURLEncoding.Decode(rawPayload, []byte(tokenSplit[1]))
	if err != nil {
		return err
	}

	var payload JWTPayload
	err = json.Unmarshal(rawPayload, &payload)
	if err != nil {
		return err
	}

	// TODO validate

	return nil
}
*/

func UnpackUserToken(token string) (user UserToken, err error) {
	return user, err
}

func UnpackApplicationToken(token string) (user UserToken, err error) {
	return user, err
}

