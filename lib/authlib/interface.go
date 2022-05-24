package authlib

import "context"

type Authenticator interface {
	GenerateServiceToken(serviceID int) string
	GenerateServiceTokenWithUserData(serviceID int, ctx context.Context) string
}

type TokenParser interface {
	ParseUserToken(token string) (UserToken, error)
	ParseApplicationToken(token string) (ApplicationToken, error)
}

type KeyProvider interface {
	GetPrivateKey() (string, error)
	GetPublicKey(applicationID int) (string, error)
}
