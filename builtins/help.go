package builtins

import (
    "fmt"
    "io"
)

// Help prints help information about available commands.
func Help(w io.Writer) error {
    helpText := `
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
  touch     - Change file timestamps or create new files.
  tail      - Display the last part of files.
  printf    - Format and print data.
  sleep     - Pause operations for a specified time.
`
    _, err := fmt.Fprint(w, helpText)  // Changed from Fprintln to Fprint to control newlines precisely
    return err
}
