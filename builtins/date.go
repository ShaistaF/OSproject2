package builtins

import (
    "bytes"
    "strings"
    "testing"
    "time"
)

func TestDate(t *testing.T) {
    w := &bytes.Buffer{}
    if err := Date(w); err != nil {
        t.Fatalf("Date() returned an error: %v", err)
    }
    got := w.String()
    expectedDate := time.Now().Format(time.RFC1123)

    if !strings.Contains(got, expectedDate) {
        t.Errorf("Date() = %q, want %q", got, expectedDate)
    }
}
