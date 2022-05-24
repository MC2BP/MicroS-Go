package middleware

import (
	"net/http"
	"strings"

	"github.com/MC2BP/MicroS-Go/lib/loglib"
	"github.com/gorilla/mux"
)

var AllowedMethods = "GET,HEAD,POST,PUT,DELETE,OPTIONS,PATCH"
var AllowedHeaders = "Content-Type,Authorization,Origin"

type basicCorsMiddleware struct {
	origins        []string
	allowedOrigins string
}

func CorsMiddleware(origins []string) mux.MiddlewareFunc {
	if len(origins) == 0 {
		origins = []string{"*"}
	}
	allowedOrigins := strings.Join(origins, ",")

	cors := &basicCorsMiddleware{
		allowedOrigins: allowedOrigins,
		origins:        origins,
	}
	return cors.Middleware
}

func (c *basicCorsMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set header and return if it is of method option
		w.Header().Set("Access-Control-Allow-Origin", c.allowedOrigins)
		w.Header().Set("Access-Control-Allowed-Methods", AllowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", AllowedHeaders)

		if r.Method == http.MethodOptions {
			return
		}

		for _, origin := range c.origins {
			if v, has := r.Header["Origin"]; has && len(v) > 0 && v[0] == origin {
				next.ServeHTTP(w, r)
				return
			}
		}
		loglib.Warningf("Request to '%s' blocked by cors. Header 'origin' with value '%s' not in list %s", r.URL, r.Header[originHeader], c.allowedOrigins)
		next.ServeHTTP(w, r)
	})
}

