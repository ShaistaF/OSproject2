
package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SetAlias sets an alias for a command.
func SetAlias(w io.Writer, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: alias name=value")
	}
	alias := strings.Join(args, " ")
	f, err := os.OpenFile(".aliases", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%s\n", alias)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "Alias set: %s\n", alias)
	return err
}
