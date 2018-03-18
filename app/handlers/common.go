package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type httpResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type data struct {
	name    string
	content interface{}
}

// respond with some data
func respond(w http.ResponseWriter, code int, data data) {
	// sanitize key names as best we can
	// this shouldn't be necessary, but in case
	// we malform a key by accident, it will help
	if strings.ContainsRune(data.name, ' ') {
		data.name = strings.ToLower(data.name)
		data.name = strings.Replace(data.name, " ", "_", -1)
	}

	resp := httpResponse{
		Code:    code,
		Message: http.StatusText(code),
		Data: map[string]interface{}{
			data.name: data.content,
		},
	}

	// marshal the response
	body, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return // log and bail in case of error
	}

	// write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintln(w, string(body))
}

func fail(w http.ResponseWriter, code int, message string) {
	resp := httpResponse{
		Code:    code,
		Message: http.StatusText(code),
		Data: map[string]interface{}{
			"error": message,
		},
	}

	// marshal the response
	body, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return // log and bail in case of error
	}

	// write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintln(w, string(body))
}
