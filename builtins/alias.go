package builtins

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// SetAlias sets an alias for a command.
func SetAlias(w io.Writer, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: alias name=value")
	}

	parts := strings.SplitN(args[0], "=", 2)
	if len(parts) != 2 {
		return errors.New("invalid alias format")
	}

	name, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

	// Open or create the alias file.
	f, err := os.OpenFile(".aliases", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open alias file: %w", err)
	}
	defer f.Close()

	// Write the alias to the file.
	if _, err := fmt.Fprintf(f, "%s=%s\n", name, value); err != nil {
		return fmt.Errorf("failed to write alias to file: %w", err)
	}

	// Confirm the alias was set.
	if _, err := fmt.Fprintf(w, "Alias set: %s=%s\n", name, value); err != nil {
		return fmt.Errorf("failed to confirm alias set: %w", err)
	}

	return nil
}
