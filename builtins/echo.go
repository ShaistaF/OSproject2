package builtins

import (
    "fmt"
    "io"
    "strings"
)

// Echo prints arguments to the standard output.
func Echo(w io.Writer, args ...string) error {
    fmt.Fprintln(w, strings.Join(args, " "))
    return nil
}
