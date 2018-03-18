package handlers

import (
	"net/http"

	"github.com/golangdorset/webapp/app"
)

// RegisterRoutes gets called whenever the server starts
func RegisterRoutes(server *app.Server) {
	//default handlers
	server.Router.NotFoundHandler = NotFound(server)

	// get
	server.Router.HandleFunc("/", Root(server)).Methods(http.MethodGet)
	server.Router.HandleFunc("/greet/{name}", Greeter(server)).Methods(http.MethodGet)

	// post
	server.Router.HandleFunc("/echo", Echo(server)).Methods(http.MethodPost)
}
