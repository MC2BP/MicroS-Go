package webserver

import (
	"github.com/gorilla/mux"
)

type Handler interface {
	AddRoutes(r *mux.Router)
}

