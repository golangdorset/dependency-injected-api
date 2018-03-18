package main

import (
	"log"

	"github.com/golangdorset/webapp/app"
	"github.com/golangdorset/webapp/app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// set some options
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// create the server
	// todo: pass in the fake cache implementation
	server := app.New(mux.NewRouter())

	// register the routes
	handlers.RegisterRoutes(server)

	// start
	server.Start()
}
