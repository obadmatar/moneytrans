package server

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

// encode writes data to the http response as JSON-encoded
// and sets the Content-Type header to "application/json"
func (s *server) encode(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// encode data to json
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// add headers
	for h, v := range headers {
		w.Header()[h] = v
	}

	// set response status and content-type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}