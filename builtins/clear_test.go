package builtins

import (
    "bytes"
    "testing"
)

func TestClear(t *testing.T) {
    w := &bytes.Buffer{}
    if err := Clear(w); err != nil {
        t.Fatalf("Clear() returned an error: %v", err)
    }
    expected := "\033[H\033[2J"
    if got := w.String(); got != expected {
        t.Errorf("Clear() = %q, want %q", got, expected)
    }
}

