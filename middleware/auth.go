package middleware

import (
	"net/http"
	"regexp"

	"github.com/MC2BP/MicroS-Go/lib/authlib"
	"github.com/MC2BP/MicroS-Go/lib/configlib"
	"github.com/MC2BP/MicroS-Go/lib/contextlib"
	"github.com/MC2BP/MicroS-Go/lib/loglib"
	"github.com/gorilla/mux"
)

type basicAuthMiddleware struct {
	applicationID  int
	parser         authlib.TokenParser
	excludedRoutes []*regexp.Regexp
}

func AuthMiddleware(cfg configlib.Configer, parser authlib.TokenParser, excludedRoutes []string) mux.MiddlewareFunc {
	routeRegexes := make([]*regexp.Regexp, len(excludedRoutes))
	for i, route := range excludedRoutes {
		routeRegex, err := regexp.Compile(route)
		if err != nil {
			panic(err)
		}
		routeRegexes[i] = routeRegex
	}

	auth := &basicAuthMiddleware{
		applicationID: cfg.GetApplicationID(),
		parser:        parser,
		excludedRoutes: routeRegexes,
	}
	return auth.Middleware
}

func (a *basicAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contextlib.NewContext(r)
		
		// check if routes is excluded
		for _, routeRegex := range a.excludedRoutes {
			if routeRegex.Match([]byte(r.URL.RequestURI())) {
				next.ServeHTTP(w, r)
				return
			}
		}


		var userToken string
		if v, has := r.Header["Authorization"]; has && len(v) == 1 {
			appToken, err := a.parser.ParseApplicationToken(v[0])
			if err != nil {
				loglib.Warning(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if appToken.ApplicationID != a.applicationID {
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
		ctx.Set(authlib.KeyToken, userToken)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func setServiceValues(ctx *contextlib.Context, appToken authlib.ApplicationToken) {
	ctx.Set(authlib.KeyApplicationID, appToken.SrcApplicationID)
}

func setUserValues(ctx *contextlib.Context, userToken authlib.UserToken) {
	ctx.Set(authlib.KeyUserUID, userToken.UserUID)
	ctx.Set(authlib.KeyUserName, userToken.UserName)
	ctx.Set(authlib.KeyEmail, userToken.Email)
	ctx.Set(authlib.KeyRoles, userToken.Roles)
	ctx.Set(authlib.KeyPermission, userToken.Permissions)
	ctx.Set(authlib.KeyValidUntil, userToken.ValidUntil)
}
