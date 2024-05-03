package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Export sets environment variables.
func Export(w io.Writer, args ...string) error {
	if len(args) == 0 {
		// If no arguments provided, print all environment variables
		for _, env := range os.Environ() {
			fmt.Fprintln(w, env)
		}
		return nil
	}

	// Handle setting environment variables
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid argument format: %s", arg)
		}
		os.Setenv(parts[0], parts[1])
	}
	return nil
}
