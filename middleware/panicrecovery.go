package middleware

import (
	"net/http"

	"github.com/MC2BP/MicroS-Go/lib/loglib"
	"github.com/gorilla/mux"
)

func RecoverPanicMiddleWare() mux.MiddlewareFunc {
	return mux.MiddlewareFunc(recoverPanic)
}

func recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// handle panic
				loglib.Errorf("Paniced on endpoint '%s', method %s: %s", r.URL.RequestURI(), r.Method, rec)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
