package authlib

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/MC2BP/MicroS-Go/lib/errorlib"
)

const (
	jwtHeader string = `{"alg":"HS512","typ":"JWT"}`
)

type UserToken struct {
	UserName    string
	Roles       []string
	Permissions map[string]string
	ValidUntil  time.Time
}

type JWTHeader struct {
	ALG string `json:"alg"`
	Typ string `json:"typ"`
}

type JWTPayload struct {
}

func DecodeToken(token string, result *interface{}) error {
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

	// validate

	return nil
}
