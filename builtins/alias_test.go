package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestSetAlias(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectedErr bool
	}{
		{"Valid Alias", []string{"ls", "ls -la"}, false},
		{"Missing Value", []string{"ls"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for testing.
			tmpFile, err := os.CreateTemp("", "aliases")
			if err != nil {
				t.Fatalf("Failed to create temporary file: %v", err)
			}
			defer os.Remove(tmpFile.Name()) // Clean up the temporary file.

			// Create a buffer to capture the output.
			var buf bytes.Buffer

			// Call the SetAlias function with the test arguments and the buffer as the writer.
			err = SetAlias(&buf, tt.args...)

			// Check if there was an error.
			if (err != nil) != tt.expectedErr {
				t.Errorf("SetAlias() error = %v, expected error = %v", err, tt.expectedErr)
				return
			}

			// Check if the alias was written to the temporary file.
			if _, err := os.Stat(tmpFile.Name()); os.IsNotExist(err) {
				t.Errorf("SetAlias() failed to write alias to file")
			}

			// Verify the output in the buffer matches the expected output.
			expectedOutput := fmt.Sprintf("Alias set: %s\n", strings.Join(tt.args, " "))
			if buf.String() != expectedOutput {
				t.Errorf("SetAlias() produced unexpected output: got %s, want %s", buf.String(), expectedOutput)
			}
		})
	}
}
