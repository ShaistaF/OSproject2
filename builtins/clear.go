package builtins

import (
    "bytes"
    "io"
)

// Clear clears the contents of the provided io.Writer and writes ANSI escape codes
// to clear the terminal screen.
func Clear(w io.Writer) error {
    // Check if the writer is a buffer and clear it if so
    if buf, ok := w.(*bytes.Buffer); ok {
        buf.Reset()
    }
    // ANSI escape codes to clear the screen
    _, err := w.Write([]byte("\033[H\033[2J"))
    return err
}

