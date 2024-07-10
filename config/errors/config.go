package errors

import "fmt"

const (
	CONFIG_SET   = "set"
	CONFIG_UNSET = "unset"
)

type Config struct {
	Operation string
	Key       string
	Err       error
}

func (c Config) Error() string {
	return fmt.Sprintf("Failed to %s config key: %s. Error: %v", c.Operation, c.Key, c.Err)
}
