package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server defines a standard server struct
// this is where dependencies should be embeded
type Server struct {
	Server *http.Server
	Router *mux.Router
}

// New returns a new server
// any new dependencies should be passed via this method
// todo: add the fake cache implementation
func New(router *mux.Router) *Server {
	return &Server{
		Server: &http.Server{
			Addr:              ":8080",
			Handler:           router,
			ReadTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       5 * time.Second,
		},
		Router: router,
	}
}

// Start the server
func (s *Server) Start() {
	log.Fatal(s.Server.ListenAndServe())
}
