package builtins

import (
    "bytes"
    "testing"
)

func TestEcho(t *testing.T) {
    w := new(bytes.Buffer)
    Echo(w, "Hello", "world!")
    if got := w.String(); got != "Hello world!\n" {
        t.Errorf("Expected 'Hello world!', got %s", got)
    }

    // Test echoing nothing
    w.Reset()
    Echo(w)
    if got := w.String(); got != "\n" {
        t.Errorf("Expected newline, got %s", got)
    }

    // Test echoing with multiple spaces
    w.Reset()
    Echo(w, "This", " ", "is", "a", "test")
    if got := w.String(); got != "This   is a test\n" {
        t.Errorf("Expected 'This   is a test', got %s", got)
    }
}
