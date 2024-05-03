package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Echo prints arguments to standard output.
func Echo(w io.Writer, args ...string) error {
	// Join arguments with spaces and print to writer
	_, err := fmt.Fprintln(w, strings.Join(args, " "))
	return err
}
