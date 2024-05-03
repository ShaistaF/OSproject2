package builtins

import (
    "bytes"
    "testing"
)

func TestAlias(t *testing.T) {
    w := new(bytes.Buffer)
    Alias(w, "ls='ls -l'")
    if got := w.String(); !strings.Contains(got, "ls='ls -l'") {
        t.Errorf("Expected alias not found, got %s", got)
    }
}
