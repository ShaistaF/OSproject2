package builtins

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Source reads and executes commands from a file in the current shell environment.
func Source(w io.Writer, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: source filename [arg1 arg2 ...]")
	}

	// Open the file for reading.
	file, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a scanner to read lines from the file.
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file.
	for scanner.Scan() {
		// Get the command from the current line.
		line := scanner.Text()

		// Skip empty lines and comments.
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Execute the command and capture its output.
		cmd := exec.Command("sh", "-c", line)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to execute command '%s': %v", line, err)
		}

		// Write the command output to the provided writer.
		if _, err := w.Write(output); err != nil {
			return fmt.Errorf("failed to write command output: %v", err)
		}
	}

	// Check for any errors encountered during scanning.
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
