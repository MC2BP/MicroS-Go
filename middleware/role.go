package middleware

import (
	"net/http"
	"regexp"

	"github.com/MC2BP/MicroS-Go/lib/contextlib"
	"github.com/gorilla/mux"
)

type basicRoleMiddleware struct {
	roles []EndpointRole
}

type EndpointRole struct {
	Endpoint string
	Role    string
	regex    *regexp.Regexp
}

func RoleMiddleware(roles []EndpointRole) mux.MiddlewareFunc {
	var err error 

	for _, role := range roles {
		role.regex, err = regexp.Compile(role.Endpoint)
		if err != nil {
			panic(err)
		}
	}
	mw := &basicRoleMiddleware{
		roles: roles,
	}
	return mw.Middleware
}

func (b *basicRoleMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contextlib.NewContext(r)
		roles := ctx.GetRoles()

		for _, route := range b.roles {
			if route.regex.Match([]byte(r.URL.RequestURI())) {
				if roles.Has(route.Role) {
					next.ServeHTTP(w, r)
				} else {
					w.WriteHeader(http.StatusForbidden)
				}
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
