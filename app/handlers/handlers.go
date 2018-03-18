package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/golangdorset/webapp/app"
)

// NotFound is our default not found handler
func NotFound(server *app.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fail(w, http.StatusNotFound, "the requested resource could not be found")
	}
}

// Root is the default endpoint for our web application
func Root(server *app.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// we will never use this elsewhere, so we define it inside the method
		type appRoute struct {
			Path   string `json:"path"`
			Method string `json:"method"`
		}
		var endpoints []appRoute

		// this walk method has a set signature, tho we do not need all of it
		// in order to avoid allocating variables for feeds we don't use, we use the bitbucket (_) instead
		err := server.Router.Walk(func(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
			path, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			methods, err := route.GetMethods()
			if err != nil {
				return err
			}
			ro := appRoute{
				Path:   path,
				Method: strings.Join(methods, ", "),
			}
			endpoints = append(endpoints, ro)

			return nil
		})
		if err != nil {
			fail(w, http.StatusInternalServerError, "cannot walk router")
			return
		}
		respond(w, http.StatusOK, data{
			name:    "routes",
			content: endpoints,
		})
	}
}

// Greeter will greet the user with the name he provided
func Greeter(_ *app.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]

		respond(w, http.StatusOK, data{
			name:    "greeting",
			content: fmt.Sprintf("hello %s", name),
		})
	}
}

// Echo takes in a specific JSON payload
// and sends it back
func Echo(_ *app.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fail(w, http.StatusInternalServerError, "cannot read request body")
			return
		}
		defer r.Body.Close()

		var input echoInput

		if err := json.Unmarshal(body, &input); err != nil {
			fail(w, http.StatusUnprocessableEntity, "cannot read request body")
			return
		}

		respond(w, http.StatusOK, data{
			name:    "echo",
			content: input,
		})
	}
}

type echoInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
