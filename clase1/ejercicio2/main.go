package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type greetingRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	r := chi.NewRouter()

	// POST a greeting message
	r.Post("/greeting", func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body into a new `Request` struct
		var req greetingRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Format a response string and send it as the response
		res := fmt.Sprintf("Hello %s %s", req.FirstName, req.LastName)
		w.Write([]byte(res))
	})

	http.ListenAndServe(":8080", r)
}
