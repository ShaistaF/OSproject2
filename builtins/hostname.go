package builtins

import (
    "fmt"
    "io"
    "os"
)

// Hostname prints the system's hostname.
func Hostname(w io.Writer) error {
    hostname, err := os.Hostname()
    if err != nil {
        return err
    }
    _, err = fmt.Fprintln(w, hostname)
    return err
}
