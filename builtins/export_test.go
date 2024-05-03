package builtins

import (
    "bytes"
    "os"
    "testing"
)

func TestExport(t *testing.T) {
    w := new(bytes.Buffer)
    // Test setting an environment variable
    Export(w, "TEST_VAR=hello_world")
    if val, exists := os.LookupEnv("TEST_VAR"); !exists || val != "hello_world" {
        t.Errorf("Expected environment variable 'TEST_VAR' to be 'hello_world', got %s", val)
    }

    // Test listing all environment variables
    w.Reset()
    Export(w)
    if got := w.String(); !strings.Contains(got, "TEST_VAR=hello_world") {
        t.Errorf("Expected listing to include 'TEST_VAR=hello_world', got %s", got)
    }

    // Test invalid argument
    w.Reset()
    Export(w, "INVALID")
    if got := w.String(); !strings.Contains(got, "invalid argument") {
        t.Errorf("Expected error message for invalid argument, got %s", got)
    }
}
