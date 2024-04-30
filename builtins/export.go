
package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Export sets environment variables.
func Export(w io.Writer, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: export name=value")
	}
	parts := strings.SplitN(args[0], "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid argument: %s", args[0])
	}
	key, value := parts[0], parts[1]
	os.Setenv(key, value)
	_, err := fmt.Fprintf(w, "Exported variable: %s=%s\n", key, value)
	return err
}
