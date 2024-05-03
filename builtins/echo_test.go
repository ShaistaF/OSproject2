package builtins_test

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		wantOutput    string
		expectedError error
	}{
		{
			name:       "Echo one argument",
			args:       []string{"Hello, World!"},
			wantOutput: "Hello, World!\n",
		},
		{
			name:       "Echo multiple arguments",
			args:       []string{"Hello", "World!"},
			wantOutput: "Hello World!\n",
		},
		{
			name:       "Echo empty arguments",
			args:       []string{},
			wantOutput: "\n",
		},
		{
			name:       "Echo arguments with spaces",
			args:       []string{"Hello", "   ", "World!"},
			wantOutput: "Hello    World!\n",
		},
		{
			name:          "Echo with write error",
			args:          []string{"Hello, World!"},
			wantOutput:    "",
			expectedError: errors.New("write error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			// Simulate write error if expected
			if tt.expectedError != nil {
				w := &errorWriter{&buf, tt.expectedError}
				if err := builtins.Echo(w, tt.args...); err == nil {
					t.Fatalf("Echo() expected error, got none")
				}
				return
			}

			// Call Echo with the test arguments
			if err := builtins.Echo(&buf, tt.args...); err != nil {
				t.Fatalf("Echo() unexpected error: %v", err)
			}

			// Check if the output matches the expected output
			gotOutput := buf.String()
			if gotOutput != tt.wantOutput {
				t.Errorf("Echo() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

// errorWriter writes to buffer but returns an error on Write
type errorWriter struct {
	io.Writer
	err error
}

func (ew *errorWriter) Write(p []byte) (n int, err error) {
	return 0, ew.err
}
