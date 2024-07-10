package errors

import (
	"fmt"
	"net/http"
)

// EntityNotFound is used when a given entity is not found in the system
type EntityNotFound struct {
	ID   string
	Key  string
	Name string
}

func (e EntityNotFound) Error() string {
	if e.Key == "" {
		e.Key = "ID"
	}

	if e.Name == "" {
		e.Name = "entity"
	}

	return fmt.Sprintf("no '%s' found for %s: '%s'", e.Name, e.Key, e.ID)
}

func (e EntityNotFound) StatusCode() int {
	return http.StatusNotFound
}

func (e EntityNotFound) EncapsulateError() map[string]any {
	return map[string]any{"error": e.Error()}
}
