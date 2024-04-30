
package builtins

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)


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
		err := handleInput(w, line, nil) // Assuming nil for 'exit' channel
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
