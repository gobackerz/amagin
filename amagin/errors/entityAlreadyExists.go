package errors

import "fmt"

type EntityAlreadyExists struct {
	Name string
}

func (e EntityAlreadyExists) Error() string {
	return fmt.Sprintf("Entity '%s' already exists", e.Name)
}
