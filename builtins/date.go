// date.go
package builtins

import (
    "io"
    "time"
    "fmt"
)


func Date(w io.Writer) error {
    now := time.Now()
    _, err := fmt.Fprintf(w, now.Format("2006-01-02")) 
    return err
}
