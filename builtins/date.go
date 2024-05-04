package builtins

import (
    "io"
    "time"
)

const DefaultDateFormat = time.RFC1123 


func Date(w io.Writer, format string) error {
    if format == "" {
        format = DefaultDateFormat
    }
    currentTime := time.Now().Format(format)
    _, err := w.Write([]byte(currentTime))
    return err
}
