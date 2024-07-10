package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type EntityAlreadyExists struct {
	Name              string
	ConflictingFields []string
}

func (e EntityAlreadyExists) Error() string {
	if e.Name == "" {
		e.Name = "entity"
	}

	conflictingFields := strings.TrimSuffix(strings.Join(e.ConflictingFields, ", "), ", ")

	var errMessage string

	switch len(e.ConflictingFields) {
	case 0:
		errMessage = fmt.Sprintf("%s with given combination already exists", e.Name)
	case 1:
		errMessage = fmt.Sprintf("%s with given %s already exists", e.Name, e.ConflictingFields[0])
	default:
		errMessage = fmt.Sprintf("%s with given %s already exists", e.Name, conflictingFields)
	}

	return errMessage
}

func (e EntityAlreadyExists) StatusCode() int {
	return http.StatusConflict
}

func (e EntityAlreadyExists) EncapsulateError() map[string]any {
	return map[string]any{"error": e.Error()}
}
