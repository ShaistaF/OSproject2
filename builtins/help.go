package builtins

import (
    "fmt"
    "io"
)


func Help(w io.Writer) error {
    helpText := `
Available Commands:
  help      - Display this help message.
  echo      - Echo the input arguments.
  cd        - Change the directory.
  env       - List all environment variables.
  pwd       - Display the current working directory.
  sleep     - Pause operations for a specified time.
  test      - Evaluate conditional expressions and return true or false.
`
    _, err := fmt.Fprint(w, helpText)  
    return err
}
