
package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Test evaluates conditional expressions.
func Test(w io.Writer, args ...string) error {
	if len(args) < 3 {
		return fmt.Errorf("usage: test arg1 operator arg2")
	}
	arg1, operator, arg2 := args[0], args[1], args[2]
	result := false
	switch operator {
	case "=":
		result = arg1 == arg2
	case "!=":
		result = arg1 != arg2
	// Add other operators as needed
	default:
		return fmt.Errorf("unsupported operator: %s", operator)
	}
	if result {
		_, err := fmt.Fprintln(w, "true")
		return err
	}
	_, err := fmt.Fprintln(w, "false")
	return err
}
