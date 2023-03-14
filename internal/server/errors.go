package server

import (
	"fmt"
	"net/http"
	"os"
)

// logError generic helper for logging error messages
func (s *server) logError(r *http.Request, err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
}

// errorResponse generic helper for sending JSON-fromatted
// error message with the given status code
func (s *server) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := s.encode(w, status, env, nil)
	if err != nil {
		s.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// notFoundResponse sends 500 internal server error error message
func (s *server) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	s.logError(r, err)

	message := "the server encounterd a problem and could not process your request"
	s.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse sends 404 not found status with error message
func (s *server) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resouce could not be found"
	s.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowed sends 405 method not allowed status with error message
func (s *server) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	s.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// badRequestResponse sends 400 bad request status with error message
func (s *server) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	s.errorResponse(w, r, http.StatusBadRequest, err.Error())
}