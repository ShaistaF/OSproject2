package builtins

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

// Source executes a script within the current shell environment.
func Source(w io.Writer, args ...string) error {
    if len(args) != 1 {
        fmt.Fprintln(w, "source: expected exactly one argument")
        return nil
    }
    file, err := os.Open(args[0])
    if err != nil {
        fmt.Fprintf(w, "source: cannot open %s: %v\n", args[0], err)
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Fprintln(w, line) // Simulating execution; ideally, you should execute this line.
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintf(w, "source: error reading file: %v\n", err)
        return err
    }
    return nil
}
