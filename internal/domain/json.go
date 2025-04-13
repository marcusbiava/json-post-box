package domain

import (
	"fmt"
)

// JSON represents a JSON document entity
type JSON struct {
	ID   string      `json:"id"`
	Data interface{} `json:"data"`
}

// ErrInvalidJSON is returned when the JSON data is invalid
var ErrInvalidJSON = fmt.Errorf("invalid JSON data")

// JSONRepository defines the interface for JSON storage operations
type JSONRepository interface {
	Store(data *JSON) error
	FindByID(id string) (*JSON, error)
}

// Validate checks if the JSON data is valid
func (j *JSON) Validate() error {
	if j.Data == nil {
		return ErrInvalidJSON
	}
	return nil
}
