package builtins

import (
    "bytes"
    "testing"
)

func TestClear(t *testing.T) {
    // Test 1: Standard case where buffer should be cleared and contain the ANSI clear code
    w := &bytes.Buffer{}
    if err := Clear(w); err != nil {
        t.Fatalf("Clear() returned an error: %v", err)
    }
    expected := "\033[H\033[2J"
    if got := w.String(); got != expected {
        t.Errorf("Clear() = %q, want %q", got, expected)
    }

    // Test 2: Ensure Clear does not append to existing content (should overwrite)
    w.Reset() // Clear the buffer from previous test
    w.WriteString("existing content")
    if err := Clear(w); err != nil {
        t.Fatalf("Clear() returned an error after writing existing content: %v", err)
    }
    if got := w.String(); got != expected {
        t.Errorf("After writing to buffer, Clear() = %q, want %q", got, expected)
    }

    // Test 3: Buffer with only whitespace before calling Clear
    w.Reset() // Clear the buffer from previous test
    w.WriteString("     ") // Write only whitespace
    if err := Clear(w); err != nil {
        t.Fatalf("Clear() returned an error after writing whitespace: %v", err)
    }
    if got := w.String(); got != expected {
        t.Errorf("After whitespace in buffer, Clear() = %q, want %q", got, expected)
    }

    // Test 4: Empty buffer case
    w.Reset() // Ensure the buffer is empty
    if err := Clear(w); err != nil {
        t.Fatalf("Clear() returned an error on empty buffer: %v", err)
    }
    if got := w.String(); got != expected {
        t.Errorf("On empty buffer, Clear() = %q, want %q", got, expected)
    }
}
