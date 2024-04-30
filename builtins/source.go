// Package builtins provides implementations of shell built-in commands.
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			// Skip empty lines and comments.
			continue
		}
		
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
