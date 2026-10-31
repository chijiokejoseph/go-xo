package utils

import "fmt"

func WrapErr(err, wrapper error) error {
	return fmt.Errorf("Error: %w, Details: %s", wrapper, err)
}
