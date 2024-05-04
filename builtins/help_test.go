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
  version   - Show the application version.
  hostname  - Display the system's hostname.
  help      - Display this help message.
  echo      - Echo the input arguments.
  exit      - Exit the application.
  cd        - Change the directory.
  env       - List all environment variables.
  clear     - Clear the terminal screen.
  date      - Display the current date and time.
  whoami    - Show the current system user.
  pwd       - Display the current working directory.
  sleep     - Pause operations for a specified time.
`
    if got := w.String(); got != expected {
        t.Errorf("Help() = %q, want %q", got, expected)
    }
}
