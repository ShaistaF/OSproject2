package builtins

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Test evaluates conditional expressions.
func Test(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no test expression provided")
	}

	cmd := exec.Command("test", args...)
	cmd.Stdout = w
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("test failed: %v", err)
	}

	return nil
}
