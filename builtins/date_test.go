package builtins

import (
    "bytes"
    "testing"
    "time"
)

func TestDateDefaultFormat(t *testing.T) {
    w := &bytes.Buffer{}
    if err := Date(w, ""); err != nil {
        t.Fatalf("Date() returned an error: %v", err)
    }
    expected := time.Now().Format(DefaultDateFormat)
    if got := w.String(); got != expected {
        t.Errorf("Date() = %q, want %q", got, expected)
    }
}

func TestDateCustomFormat(t *testing.T) {
    w := &bytes.Buffer{}
    customFormat := "2006-01-02"
    if err := Date(w, customFormat); err != nil {
        t.Fatalf("Date(customFormat) returned an error: %v", err)
    }
    expected := time.Now().Format(customFormat)
    if got := w.String(); got != expected {
        t.Errorf("Date(customFormat) = %q, want %q", got, expected)
    }
}

func TestDateEmptyWriter(t *testing.T) {
    w := &bytes.Buffer{} // Simulate an empty buffer
    if err := Date(w, DefaultDateFormat); err != nil {
        t.Fatalf("Date() with empty writer returned an error: %v", err)
    }
    if w.Len() == 0 {
        t.Error("Date() did not write anything to buffer, expected date string")
    }
}
