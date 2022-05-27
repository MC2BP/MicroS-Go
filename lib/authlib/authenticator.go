package authlib

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"strings"
	"sync"

	"context"

	"github.com/MC2BP/MicroS-Go/lib/configlib"
	"github.com/MC2BP/MicroS-Go/lib/errorlib"
)

const (
	KeyUserName      = "uname"
	KeyUserUID       = "uid"
	KeyEmail         = "email"
	KeyRoles         = "roles"
	KeyPermission    = "permission"
	KeyValidUntil    = "validuntil"
	KeyApplicationID = "applicationid"
	KeyToken         = "token"
	
	UserAuthServiceID = 1
)

var defaultJWTHeader string

func init() {
	header, err := json.Marshal(JWTHeader{
		Alg: "RS512",
		Typ: "JWT",
	})
	if err != nil {
		panic(err)
	}
	defaultJWTHeader = base64.RawURLEncoding.EncodeToString(header)
}

type basicAuthenticator struct {
	keyProvider KeyProvider
	privateKey  *rsa.PrivateKey
	publicKeys  map[int]*rsa.PublicKey
	appID       int
	sync.Mutex
}

func NewBasicAuthenticator(cfg configlib.Configer, keyProvider KeyProvider) *basicAuthenticator {
	key, err := keyProvider.GetPrivateKey()
	if err != nil {
		panic(err)
	}
	return &basicAuthenticator{
		privateKey: key,
		appID:      cfg.GetApplicationID(),
	}
}

func (a *basicAuthenticator) GenerateServiceToken(serviceID int) (string, error) {
	payloadRaw, err := json.Marshal(ApplicationToken{
		SrcApplicationID: a.appID,
		ApplicationID:    serviceID,
	})
	if err != nil {
		return "", err
	}

	payload := base64.RawURLEncoding.EncodeToString(payloadRaw)
	return a.getJWT(serviceID, defaultJWTHeader, payload)
}

func (a *basicAuthenticator) GenerateServiceTokenWithUserData(serviceID int, ctx context.Context) (string, error) {
	payloadRaw, err := json.Marshal(ApplicationToken{
		SrcApplicationID: a.appID,
		ApplicationID:    serviceID,
		UserToken:        ctx.Value(KeyToken).(string),
	})
	if err != nil {
		return "", err
	}

	payload := base64.RawURLEncoding.EncodeToString(payloadRaw)
	return a.getJWT(serviceID, defaultJWTHeader, payload)
}

func (a *basicAuthenticator) ParseUserToken(token string) (payload UserToken, err error) {
	jwt := strings.Split(token, ".")
	if len(jwt) != 3 {
		return payload, errorlib.Err("JWT invalid")
	}
	
	rawPayload, err := base64.RawStdEncoding.DecodeString(jwt[1])
	if err != nil {
		return
	}
	err = json.Unmarshal(rawPayload, &payload)
	if err != nil {
		return 
	}

	err = a.verifySignature(UserAuthServiceID, jwt[0], jwt[1], jwt[2])
	return 
}

func (a *basicAuthenticator) ParseApplicationToken(token string) (ApplicationToken, error) {
	return ApplicationToken{}, nil
}

func (a *basicAuthenticator) GetPublicKey(appID int) (*rsa.PublicKey, error) {
	a.Lock()
	defer a.Unlock()
	if key, ok := a.publicKeys[appID]; ok {
		return key, nil
	}
	key, err := a.keyProvider.GetPublicKey(appID)
	if err != nil {
		return nil, err
	}
	a.publicKeys[appID] = key
	return key, nil
}

func getHash(header string, payload string) []byte {
	msgHash := sha512.New()
	msgHash.Write([]byte(header + "." + payload))
	return msgHash.Sum(nil)
}

func (a *basicAuthenticator) getJWT(serviceID int, header, payload string) (string, error) {
	sign, err := rsa.SignPSS(rand.Reader, a.privateKey, crypto.SHA256, getHash(defaultJWTHeader, payload), nil)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString([]byte(header)) +
		"." +
		base64.RawStdEncoding.EncodeToString([]byte(payload)) +
		"." +
		base64.RawStdEncoding.EncodeToString(sign), nil
}

func (a *basicAuthenticator) verifySignature(serviceID int, header, payload , signature string) error {
	rawSign, err := base64.RawStdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	publicKey, err := a.GetPublicKey(serviceID)
	if err != nil {
		return err
	}
	hash := getHash(header, payload)
	return rsa.VerifyPSS(publicKey, crypto.SHA512, hash, rawSign, nil)
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
