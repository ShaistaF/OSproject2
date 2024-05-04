package builtins

import (
    "bytes"
    "testing"
)

func TestHistory(t *testing.T) {
    commandHistory = nil 
    
    AddToHistory("ls -l")
    AddToHistory("echo Hello World")
    AddToHistory("cd /")

    w := &bytes.Buffer{}
    if err := History(w); err != nil {
        t.Fatalf("History() returned an error: %v", err)
    }

    expected := "1 ls -l\n2 echo Hello World\n3 cd /\n"
    if got := w.String(); got != expected {
        t.Errorf("History() = %q, want %q", got, expected)
    }
}
