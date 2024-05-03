package builtins

import (
    "fmt"
    "io"
    "strings"
)

// Echo writes the arguments to the given io.Writer separated by spaces,
// followed by a newline.
func Echo(w io.Writer, args ...string) error {
    _, err := fmt.Fprintln(w, strings.Join(args, " "))
    return err
}

