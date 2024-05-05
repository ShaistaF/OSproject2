package builtins

import (
	"bytes"
	"os"
	"testing"
	"strings"
)

func TestCommand(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"-v", "echo"}, "echo\n"},
		{[]string{"-V", "echo"}, "/bin/echo is located at "}, // Path might vary
		{[]string{"-p", "echo", "hello"}, "hello\n"},
		{[]string{"echo", "hello"}, "hello\n"},
	}

	for _, test := range tests {
		output := &bytes.Buffer{}
		os.Stdout = output // redirect output to buffer for testing
		err := Command(test.args...)
		os.Stdout = os.Stdout // reset stdout

		if err != nil {
			t.Errorf("Command(%v) returned an error: %v", test.args, err)
		}
		if !strings.Contains(output.String(), test.expected) {
			t.Errorf("Expected output %q, got %q", test.expected, output.String())
		}
	}
}
