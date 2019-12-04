package util

import (
	"encoding/json"
	"net/http"
)

// ResponseOK sets the content type to JSON, the code to 200,
// and the body is encoded in json.
func ResponseOK(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

// ResponseError returns an error in json with the key as `error`
func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}

	json.NewEncoder(w).Encode(body)
}
