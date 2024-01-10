package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func writeResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{Message: message, Data: data})
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	writeResponse(w, statusCode, err.Error(), nil)
}
