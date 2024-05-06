package builtins

import (
    "bytes"
    "testing"
)

func TestHelp(t *testing.T) {
    w := &bytes.Buffer{}
    err := Help(w)
    if err != nil {
        t.Errorf("Help() error = %v", err)
    }
    expected := `
Available Commands:
  help      - Display this help message.
  echo      - Echo the input arguments.
  cd        - Change the directory.
  env       - List all environment variables.
  pwd       - Display the current working directory.
  sleep     - Pause operations for a specified time.
  test      - Evaluate conditional expressions and return true or false.
`
    if got := w.String(); got != expected {
        t.Errorf("Help() = %q, want %q", got, expected)
    }
}
