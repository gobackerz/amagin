package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type MissingParam struct {
	Param []string
}

func (e MissingParam) Error() string {
	var errMessage string

	switch len(e.Param) {
	case 0:
		errMessage = fmt.Sprintf("missing for parameters")
	case 1:
		errMessage = fmt.Sprintf("missing parameter: %s", e.Param[0])
	default:
		errMessage = fmt.Sprintf("missing parameters: %s", strings.TrimSuffix(strings.Join(e.Param, ", "), ", "))
	}

	return errMessage
}

func (e MissingParam) StatusCode() int {
	return http.StatusBadRequest
}

func (e MissingParam) EncapsulateError() map[string]any {
	return map[string]any{"error": e.Error()}
}
