package errors

import (
	"fmt"
	"github.com/gobackerz/amagin/constants"
	"strings"
)

// InvalidParam is used when invalid parameters are passed in a request
type InvalidParam struct {
	Param     []string
	Attribute string
}

// Error returns an error message regarding invalid parameters for a request
func (e InvalidParam) Error() string {
	if e.Attribute != constants.Type && e.Attribute != constants.Value {
		e.Attribute = constants.Value
	}

	if len(e.Param) > 1 {
		return fmt.Sprintf("Incorrect %s for parameters: %s", e.Attribute, strings.Join(e.Param, ", "))
	} else if len(e.Param) == 1 {
		return fmt.Sprintf("Incorrect %s for parameter: %s", e.Attribute, e.Param[0])
	}

	return fmt.Sprintf("This request has incorrect %s for parameters", e.Attribute)
}
