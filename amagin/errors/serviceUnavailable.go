package errors

import "fmt"

type ServiceUnavailable struct {
	Name string
}

func (e ServiceUnavailable) Error() string {
	return fmt.Sprintf("'%s' service unavailable", e.Name)
}
