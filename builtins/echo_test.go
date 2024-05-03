package builtins

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		output string
	}{
		{"Simple Test", []string{"Hello", "World"}, "Hello World\n"},
		{"Empty Args", []string{}, "\n"},
		{"Special Characters", []string{"This", "is", "a", "test", "with", "special", "characters: !@#$%^&*()"}, "This is a test with special characters: !@#$%^&*()\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture the output.
			var buf bytes.Buffer

			// Call the Echo function with the test arguments and the buffer as the writer.
			err := Echo(&buf, tt.args...)

			// Check if there was an error.
			if err != nil {
				t.Errorf("Echo() error = %v", err)
				return
			}

			// Compare the expected output with the actual output captured in the buffer.
			if got := buf.String(); got != tt.output {
				t.Errorf("Echo() = %v, want %v", got, tt.output)
			}
		})
	}
}
