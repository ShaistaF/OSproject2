package builtins

import (
    "bytes"
    "testing"
    "strings"
)

func TestAlias(t *testing.T) {
    w := new(bytes.Buffer)
    Alias(w, "ll='ls -alF'", "rm='rm -i'")
    expected := "ll='ls -alF'\nrm='rm -i'\n"
    if got := w.String(); got != expected {
        t.Errorf("Expected %q, got %q", expected, got)
    }

    // Test listing aliases with no arguments
    w.Reset()
    Alias(w)
    if got := w.String(); got != expected {
        t.Errorf("Expected to list %q, got %q", expected, got)
    }

    // Test invalid format
    w.Reset()
    Alias(w, "invalid")
    if got := w.String(); !strings.Contains(got, "invalid format") {
        t.Errorf("Expected error message for invalid format, got %q", got)
    }
}
