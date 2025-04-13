package domain

import "errors"

var (
	// ErrJSONNotFound is returned when a JSON document is not found
	ErrJSONNotFound = errors.New("json document not found")

	// ErrStorageFailure is returned when there's a failure storing the JSON
	ErrStorageFailure = errors.New("failed to store json")
)
