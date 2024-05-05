package builtins

import (
    "fmt"
    "io"
)


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
  sleep     - Pause operations for a specified time.
`
    _, err := fmt.Fprint(w, helpText)  
    return err
}
