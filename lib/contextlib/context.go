package contextlib

import (
	"context"
	"net/http"
	"time"

	"github.com/MC2BP/MicroS-Go/lib/authlib"
)

const (
	KeyUserName      = "uname"
	KeyUserUID       = "uid"
	KeyEmail         = "email"
	KeyRoles         = "roles"
	KeyPermission    = "permission"
	KeyValidUntil    = "validuntil"
	KeyApplicationID = "applicationid"
)

type Context struct {
	context context.Context
}

func NewContext(r *http.Request) *Context {
	return &Context{
		context: r.Context(),
	}
}

// Set saves a key-value pair in the context
func (ctx *Context) Set(key, val interface{}) {
	ctx.context = context.WithValue(ctx.context, key, val)
}

// GetUserUID returns the UID of the current user
func (ctx *Context) GetUserUID() (uid string) {
	uid, _ = ctx.Value(KeyUserUID).(string)
	return uid
}

// GetRoles returns the roles of the current user
func (ctx *Context) GetRoles() authlib.Roles {
	roles, _ := ctx.Value(KeyUserUID).(authlib.Roles)
	return roles
}

// GetRoles returns the permissions of the current user
func (ctx *Context) GetPermission() authlib.Permission {
	permission, _ := ctx.Value(KeyUserUID).(authlib.Permission)
	return permission
}

// GetEmail returns the email of the current user
func (ctx *Context) GetEmail() string {
	email, _ := ctx.Value(KeyEmail).(string)
	return email
}

// GetApplicationID returns the id of the application
func (ctx *Context) ApplicationID() string {
	email, _ := ctx.Value(KeyEmail).(string)
	return email
}

/* Implementation of context interface functions */

// Deadline returns the deadline event
func (ctx Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.context.Deadline()
}

// Value returns the context value according to the key given in parameter
func (ctx Context) Value(key interface{}) interface{} {
	return ctx.context.Value(key)
}

// Done is added to implement the interface
func (ctx Context) Done() <-chan struct{} {
	return ctx.context.Done()
}

// Err is added to implement the interface
func (ctx Context) Err() error {
	return ctx.context.Err()
}
