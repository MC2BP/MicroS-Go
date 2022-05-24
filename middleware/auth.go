package middleware

import (
	"net/http"

	"github.com/MC2BP/MicroS-Go/lib/authlib"
	"github.com/MC2BP/MicroS-Go/lib/configlib"
	"github.com/MC2BP/MicroS-Go/lib/contextlib"
	"github.com/MC2BP/MicroS-Go/lib/loglib"
	"github.com/gorilla/mux"
)

type basicAuthMiddleware struct {
	applicationID int
	parser        authlib.TokenParser
}

func AuthMiddleware(cfg configlib.Configer, parser authlib.TokenParser) mux.MiddlewareFunc {
	auth := &basicAuthMiddleware{
		applicationID: cfg.GetApplicationID(),
		parser:        parser,
	}
	return auth.Middleware
}

func (a *basicAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contextlib.NewContext(r)
		var userToken string
		if v, has := r.Header["Authorization"]; has && len(v) == 1 {
			appToken, err := a.parser.ParseApplicationToken(v[0])
			if err != nil {
				loglib.Warning(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			setServiceValues(ctx, appToken)
			userToken = appToken.UserToken
		} else if cookie, err := r.Cookie(""); err != nil {
			userToken = cookie.Value
		}

		if userToken != "" {
			userToken, err := a.parser.ParseUserToken(userToken)
			if err != nil {
				loglib.Warning(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			setUserValues(ctx, userToken)
		}

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func setServiceValues(ctx *contextlib.Context, appToken authlib.ApplicationToken) {
	ctx.Set(contextlib.KeyApplicationID, appToken.SrcApplicationID)
}

func setUserValues(ctx *contextlib.Context, userToken authlib.UserToken) {
	ctx.Set(contextlib.KeyUserUID, userToken.UserUID)
	ctx.Set(contextlib.KeyUserName, userToken.UserName)
	ctx.Set(contextlib.KeyEmail, userToken.Email)
	ctx.Set(contextlib.KeyRoles, userToken.Roles)
	ctx.Set(contextlib.KeyPermission, userToken.Permissions)
	ctx.Set(contextlib.KeyValidUntil, userToken.ValidUntil)
}
