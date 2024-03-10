package errors

import (
	"fmt"
	"strings"
)

type MissingParam struct {
	Param []string
}

func (e MissingParam) Error() string {
	if len(e.Param) > 1 {
		return fmt.Sprintf("Parameters %s are required for this request", strings.Join(e.Param, ", "))
	} else if len(e.Param) == 1 {
		return fmt.Sprintf("Parameter %s is required for this request", e.Param[0])
	}

	return "This request has missing parameters"
}
