package authlib

import "time"

type Roles []string

type ObjectPermission map[string]Roles

type Permission map[string]ObjectPermission

type UserToken struct {
	UserName    string     `json:"usr"`
	UserUID     string     `json:"uid"`
	Email       string     `json:"mail"`
	Roles       Roles      `json:"roles"`
	Permissions Permission `json:"permissions"`
	ValidUntil  time.Time  `json:"exp"`
}

type ApplicationToken struct {
	ApplicationID    int    `json:"aud"`
	SrcApplicationID int    `json:"iss"`
	UserToken        string `json:"usrt"`
}

type JWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWTPayload struct {
	Username   string
	Roles      []string
	Permission map[string]string
	ValidUntil time.Time
}
