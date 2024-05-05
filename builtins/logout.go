package builtins

import (
	"fmt"
	"os"
	"strconv"
)

// Logout exits the program with a given exit status.
func Logout(args ...string) error {
	status := 0 // default exit status is 0

	if len(args) > 0 {
		var err error
		status, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid exit status: %s", args[0])
		}
	}

	fmt.Printf("Exiting with status: %d\n", status)
	os.Exit(status)
	return nil // this will never be reached but is necessary to match the expected function signature
}
