package builtins

import (
    "bytes"
    "testing"
)

func TestVersion(t *testing.T) {
    w := &bytes.Buffer{}
    err := Version(w)
    if err != nil {
        t.Errorf("Version() error = %v", err)
    }
    expected := "Version 1.0.0\n"
    if got := w.String(); got != expected {
        t.Errorf("Version() = %q, want %q", got, expected)
    }
}

