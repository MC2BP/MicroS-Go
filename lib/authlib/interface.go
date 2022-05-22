package authlib

import "context"

type Authenticator interface {
	GenerateServiceToken(serviceID int64) string
	GenerateUserToken(ctx context.Context) string
}
