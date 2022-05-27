package authlib

import (
	"context"
	"crypto/rsa"
)

type Authenticator interface {
	GenerateServiceToken(serviceID int) (string, error)
	GenerateServiceTokenWithUserData(serviceID int, ctx context.Context) (string, error)
}

type TokenParser interface {
	ParseUserToken(token string) (UserToken, error)
	ParseApplicationToken(token string) (ApplicationToken, error)
}

type KeyProvider interface {
	GetPrivateKey() (*rsa.PrivateKey, error)
	GetPublicKey(applicationID int) (*rsa.PublicKey, error)
}
