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

    
    w.Reset() 
    w.WriteString("existing content")
    if err := Clear(w); err != nil {
        t.Fatalf("Clear() returned an error after writing existing content: %v", err)
    }
    if got := w.String(); got != expected {
        t.Errorf("After writing to buffer, Clear() = %q, want %q", got, expected)
    }

    
    w.Reset() 
    w.WriteString("     ") 
    if err := Clear(w); err != nil {
        t.Fatalf("Clear() returned an error after writing whitespace: %v", err)
    }
    if got := w.String(); got != expected {
        t.Errorf("After whitespace in buffer, Clear() = %q, want %q", got, expected)
    }

    
    w.Reset() 
    if err := Clear(w); err != nil {
        t.Fatalf("Clear() returned an error on empty buffer: %v", err)
    }
    if got := w.String(); got != expected {
        t.Errorf("On empty buffer, Clear() = %q, want %q", got, expected)
    }
}
