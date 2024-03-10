package errors

import (
	"fmt"
	"strings"
)

type EntityConflict struct {
	Name              string
	ConflictingFields []string
}

func (e EntityConflict) Error() string {
	if len(e.ConflictingFields) > 1 {
		return fmt.Sprintf("Conflict detected while updating '%s': Fields %s are already exists for another '%s'",
			e.Name, strings.Join(e.ConflictingFields, ", "), e.Name)
	} else if len(e.ConflictingFields) == 1 {
		return fmt.Sprintf("Confict detected while updating '%s': Field %s is already exists for another '%s'", e.Name,
			e.ConflictingFields[0], e.Name)
	}

	return fmt.Sprintf("Conflict detected while updating '%s'", e.Name)
}
