package builtins

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Source reads and executes commands from a file in the current shell environment.
func Source(w io.Writer, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: source filename")
	}
	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var errors []error

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			// Skip empty lines and comments.
			continue
		}
		// Execute the shell command from the source file.
		if err := handleInput(w, line); err != nil {
			errors = append(errors, fmt.Errorf("error executing command: %s", err.Error()))
		}
	}

	if err := scanner.Err(); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		var errMsg string
		for _, err := range errors {
			errMsg += fmt.Sprintf("%s\n", err.Error())
		}
		return fmt.Errorf("encountered %d error(s) while sourcing file:\n%s", len(errors), errMsg)
	}

	return nil
}
