package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/MC2BP/MicroS-Go/lib/contextlib"
	"github.com/gorilla/mux"
)

type timeout time.Duration

func TimeoutMiddleware(timeoutDuration time.Duration) mux.MiddlewareFunc {
	to := timeout(timeoutDuration)
	return to.Middleware
}

func (to timeout) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if to > 0 {
			ctx := contextlib.NewContext(r)
			ctxTimeOut, cancel := context.WithTimeout(ctx, time.Duration(to))
			defer cancel()

			r = r.WithContext(ctxTimeOut)
		}

		next.ServeHTTP(w, r)
	})
}
