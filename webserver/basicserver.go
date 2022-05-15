package webserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MC2BP/MicroS-Go/lib/configlib"
	"github.com/gorilla/mux"
)

type basicServer struct {
	server *http.Server
}

func NewServer(conf configlib.Webserver, middlewares []mux.MiddlewareFunc, handlers []Handler) *basicServer {
	router := mux.NewRouter()
	for _, h := range handlers {
		h.AddRoutes(router)
	}
	for _, middleware := range middlewares {
		router.Use(middleware)
	}

	server := &http.Server{
		Handler: router,
		Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		ReadTimeout: time.Duration(conf.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
	}

	
	return &basicServer{
		server: server,
	}
}

func (s *basicServer) Serve() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
