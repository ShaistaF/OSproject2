package builtins

import (
	"bytes"
	"os"
	"testing"
	"strings"
)

func TestExport(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectedErr bool
	}{
		{"Valid Export", []string{"TEST_VAR=test_value"}, false},
		{"Missing Value", []string{"TEST_VAR"}, true},
		{"Invalid Argument", []string{"invalid_argument"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the Export function with the test arguments.
			var buf bytes.Buffer
			err := Export(&buf, tt.args...)

			// Check if there was an error.
			if (err != nil) != tt.expectedErr {
				t.Errorf("Export() error = %v, expected error = %v", err, tt.expectedErr)
				return
			}

			// Check if the environment variable was set.
			if tt.expectedErr == false {
				envValue := os.Getenv(strings.Split(tt.args[0], "=")[0])
				expectedValue := strings.Split(tt.args[0], "=")[1]
				if envValue != expectedValue {
					t.Errorf("Export() failed to set environment variable: got %s, want %s", envValue, expectedValue)
				}
			}
		})
	}
}
