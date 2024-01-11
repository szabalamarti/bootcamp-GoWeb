package handler

import (
	"errors"
	"net/http"
	"os"
)

var ErrUnauthorized = errors.New("unauthorized")

func Authorize(r *http.Request, w http.ResponseWriter) error {
	token := r.Header.Get("Token")
	if token != os.Getenv("Token") {
		return ErrUnauthorized
	}
	return nil
}
