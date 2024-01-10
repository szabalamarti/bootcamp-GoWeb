package response

import "net/http"

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	defaultStatusCode := http.StatusInternalServerError
	if statusCode > 299 && statusCode < 600 {
		defaultStatusCode = statusCode
	}
	WriteResponseJSON(w, defaultStatusCode, err.Error(), nil)
}
