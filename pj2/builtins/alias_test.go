package builtins

import (
	"bytes"
	"testing"
)

func TestSetAlias(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedErr    bool
		expectedOutput string // Add expected output for the successful case
	}{
		// Providing correct format "name=value"
		{"Valid Alias", []string{"alias=ls -la"}, false, "Alias set: alias=ls -la\n"},
		// Incorrect format, expecting error
		{"Missing Value", []string{"ls"}, true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture the output.
			var buf bytes.Buffer

			// Call the SetAlias function with the test arguments and the buffer as the writer.
			err := SetAlias(&buf, tt.args...)

			// Check if there was an error.
			if (err != nil) != tt.expectedErr {
				t.Errorf("SetAlias() error = %v, expected error = %v", err, tt.expectedErr)
				return
			}

			// Check the output only if no error is expected
			if !tt.expectedErr && buf.String() != tt.expectedOutput {
				t.Errorf("SetAlias() produced unexpected output: got %s, want %s", buf.String(), tt.expectedOutput)
			}
		})
	}
}
