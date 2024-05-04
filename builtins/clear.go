package builtins

import (
    "io"
)

// Clear sends the ANSI escape sequence to clear the terminal screen.
func Clear(w io.Writer) error {
    _, err := w.Write([]byte("\033[H\033[2J"))
    return err
}
