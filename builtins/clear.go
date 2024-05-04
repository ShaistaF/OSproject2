package builtins

import (
    "bytes"
    "io"
)


func Clear(w io.Writer) error {
    
    if buf, ok := w.(*bytes.Buffer); ok {
        buf.Reset()
    }
    
    _, err := w.Write([]byte("\033[H\033[2J"))
    return err
}

