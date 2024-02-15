package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router instance.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	// Additional router setup can be done here
	return router
}

// Run starts the HTTP server with the given address.
func Run(router *mux.Router, addr string) error {
	return http.ListenAndServe(addr, router)
}
