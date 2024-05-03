package builtins

import (
    "fmt"
    "io"
    "sync"
)

var (
    aliasLock sync.RWMutex
    aliases   = make(map[string]string)
)

// Alias sets or displays aliases.
func Alias(w io.Writer, args ...string) error {
    aliasLock.Lock()
    defer aliasLock.Unlock()

    if len(args) == 0 {
        for alias, command := range aliases {
            fmt.Fprintf(w, "%s='%s'\n", alias, command)
        }
        return nil
    }
    for _, arg := range args {
        parts := strings.SplitN(arg, "=", 2)
        if len(parts) != 2 {
            fmt.Fprintf(w, "alias: invalid format %s\n", arg)
            continue
        }
        aliases[parts[0]] = parts[1]
    }
    return nil
}
