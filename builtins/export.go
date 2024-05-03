package builtins

import (
    "fmt"
    "io"
    "os"
    "strings"
)

// Export sets or displays environment variables.
func Export(w io.Writer, args ...string) error {
    if len(args) == 0 {
        for _, env := range os.Environ() {
            fmt.Fprintln(w, env)
        }
        return nil
    }
    for _, arg := range args {
        parts := strings.SplitN(arg, "=", 2)
        if len(parts) != 2 {
            fmt.Fprintf(w, "export: invalid argument %s\n", arg)
            continue
        }
        if err := os.Setenv(parts[0], parts[1]); err != nil {
            fmt.Fprintf(w, "export: failed to set %s\n", arg)
            return err
        }
    }
    return nil
}
