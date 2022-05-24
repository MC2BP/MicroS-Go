package authlib

import "time"

type Roles []string

type ObjectPermission map[string]Roles

type Permission map[string]ObjectPermission

type UserToken struct {
	UserName    string
	UserUID     string
	Email       string
	Roles       Roles
	Permissions Permission
	ValidUntil  time.Time
}

type ApplicationToken struct {
	ApplicationID    int
	SrcApplicationID int
	UserToken        string
}

type JWTHeader struct {
	ALG string `json:"alg"`
	Typ string `json:"typ"`
}

type JWTPayload struct {
	Username   string
	Roles      []string
	Permission map[string]string
	ValidUntil time.Time
}
