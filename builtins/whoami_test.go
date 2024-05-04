package builtins

import (
    "bytes"
    "os/user"
    "testing"
)

func TestWhoami(t *testing.T) {
    currentUser, err := user.Current()
    if err != nil {
        t.Fatalf("Failed to get current user: %v", err)
    }
    expectedUsername := currentUser.Username

    w := &bytes.Buffer{}
    if err := Whoami(w); err != nil {
        t.Fatalf("Whoami() returned an error: %v", err)
    }
    if got := w.String(); got != expectedUsername+"\n" {
        t.Errorf("Whoami() = %q, want %q", got, expectedUsername+"\n")
    }
}
