package builtins

import (
    "fmt"
    "io"
)

func Version(w io.Writer) error {
    _, err := fmt.Fprintln(w, "Version 1.0.0")
    return err
}
