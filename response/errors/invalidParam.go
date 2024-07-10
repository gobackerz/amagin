package errors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gobackerz/amagin/constants"
)

// InvalidParam is used when invalid parameters are passed in a request
type InvalidParam struct {
	Param     []string
	Attribute string
}

func (e InvalidParam) Error() string {
	if e.Attribute != constants.Type && e.Attribute != constants.Value {
		e.Attribute = constants.Value
	}

	var errMessage string

	switch len(e.Param) {
	case 0:
		errMessage = fmt.Sprintf("incorrect %s for parameters", e.Attribute)
	case 1:
		errMessage = fmt.Sprintf("incorrect %s for parameter: %s", e.Attribute, e.Param[0])
	default:
		errMessage = fmt.Sprintf("incorrect %s for parameters: %s", e.Attribute, strings.TrimSuffix(strings.Join(e.Param, ", "), ", "))
	}

	return errMessage
}

func (e InvalidParam) StatusCode() int {
	return http.StatusBadRequest
}

func (e InvalidParam) EncapsulateError() map[string]any {
	return map[string]any{"error": e.Error()}
}
