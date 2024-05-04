package builtins

import (
    "io"
    "time"
)

const DefaultDateFormat = time.RFC1123 // Default date format if none is provided

// Date writes the current date and time to the provided io.Writer in the specified format.
func Date(w io.Writer, format string) error {
    if format == "" {
        format = DefaultDateFormat
    }
    currentTime := time.Now().Format(format)
    _, err := w.Write([]byte(currentTime))
    return err
}
