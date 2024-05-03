package builtins

import (
    "fmt"
    "io"
    "time"
)

// Date prints the current date and time to the provided io.Writer.
func Date(w io.Writer) error {
    now := time.Now()
    _, err := fmt.Fprintln(w, now.Format(time.RFC1123))
    return err
}
