package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// decode parse JSON-encoded request body and store it in v
// it returnes error if unknown fields found, body limit exceeded 1MB
// or body contains invalid JOSN sysntax, invalid JOSN type or invalid field type
func (s *server) decode(w http.ResponseWriter, r *http.Request, v any) error {

	// limit request body to 1MB.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// init JSON decoder
	decoder := json.NewDecoder(r.Body)

	// only fields defined in v
	decoder.DisallowUnknownFields()

	// decode body input and store it in v
	err := decoder.Decode(&v)
	if err == nil {

		// check if body contains only one single JSON vlaue
		err = decoder.Decode(&struct{}{})
		if err != io.EOF {
			return errors.New("body must only contain a single JSON value")
		}

		return nil
	}

	var maxBytesError *http.MaxBytesError
	var syntaxtError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var invalidUmmarshalError *json.InvalidUnmarshalError

	// check if it is invalid destination
	if errors.As(err, &invalidUmmarshalError) {
		panic(err)
	}

	// check if it is empty body error
	if errors.Is(err, io.EOF) {
		return errors.New("body must not be empty")
	}

	// check if it is unexpected syntax errors in the JSON
	// open issue: https://github.com/golang/go/issues/25956
	if errors.Is(err, io.ErrUnexpectedEOF) {
		return errors.New("body contains badly-formed JSON")
	}

	// check if it is body size error
	if errors.As(err, &maxBytesError) {
		return fmt.Errorf("body must not exceed %d bytes", maxBytesError.Limit)
	}

	// check if it is unknown field error
	if strings.HasPrefix(err.Error(), "json: unknown field ") {
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		return fmt.Errorf("body contains unknown keys %v", fieldName)
	}

	// check if it is invalid syntax error
	if errors.As(err, &syntaxtError) {
		return errors.New("body contains badly-formed JSON and can not be parsed")
	}

	// check if it is invalid type error
	if errors.As(err, &unmarshalTypeError) {
		if unmarshalTypeError.Field != "" {
			return fmt.Errorf("body contains incorrect JSON type for field  %q", unmarshalTypeError.Field)
		}
		return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
	}

	return err
}