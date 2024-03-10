package errors

import "fmt"

// EntityNotFound is used when a given entity is not found in the system
type EntityNotFound struct {
	ID   string
	Name string
}

// Error returns an error message indicating that the entity is not found
func (e EntityNotFound) Error() string {
	return fmt.Sprintf("No '%s' found for ID: '%s'", e.Name, e.ID)
}
