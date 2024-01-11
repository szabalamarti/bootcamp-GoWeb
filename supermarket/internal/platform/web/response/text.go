package response

import "net/http"

// Text writes text response
func Text(w http.ResponseWriter, statusCode int, body string) {
	// set header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// set status code
	w.WriteHeader(statusCode)

	// write body
	w.Write([]byte(body))
}
