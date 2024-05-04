package builtins

import (
    "bytes"
    "os"
    "testing"
)

func TestHostname(t *testing.T) {
    w := &bytes.Buffer{}
    err := Hostname(w)
    if err != nil {
        t.Errorf("Hostname() error = %v", err)
    }
    expectedHostname, err := os.Hostname()
    if err != nil {
        t.Fatalf("Getting system hostname failed: %v", err)
    }
    expected := expectedHostname + "\n"
    if got := w.String(); got != expected {
        t.Errorf("Hostname() = %q, want %q", got, expected)
    }
}
