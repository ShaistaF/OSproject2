package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	// Join the arguments into a single string separated by spaces.
	echoOutput := strings.Join(args, " ")

	// Print the output to the provided writer.
	_, err := fmt.Fprintln(w, echoOutput)
	return err
}

