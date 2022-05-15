package example

import (
	"fmt"
	"net/http"

	"github.com/MC2BP/MicroS-Go/lib/configlib"
	"github.com/MC2BP/MicroS-Go/lib/loglib"
	"github.com/MC2BP/MicroS-Go/middleware"
	"github.com/MC2BP/MicroS-Go/webserver"
	"github.com/gorilla/mux"
)

func Main() {
	config := configlib.ReadConfig()

	middleware := []mux.MiddlewareFunc{
		middleware.RecoverPanicMiddleWare(),
		middleware.CorsMiddleware(config.GetWebserver().Cors),
	}

	handlers := []webserver.Handler{test{}}

	server := webserver.NewServer(config.GetWebserver(), middleware, handlers)
	err := server.Serve()
	if err != nil {
		loglib.Error(err)
	}
}

type test struct {}

func (t test) AddRoutes(r *mux.Router) {
	r.HandleFunc("/hello", t.Test)
}

func (t test) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
