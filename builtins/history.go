package builtins

import (
    "fmt"
    "io"
    "sync"
)

// commandHistory stores the history of commands executed.
var commandHistory []string
var lock sync.Mutex

// AddToHistory adds a command to the history.
func AddToHistory(command string) {
    lock.Lock()
    defer lock.Unlock()
    commandHistory = append(commandHistory, command)
}

// History prints the command history to the provided io.Writer.
func History(w io.Writer) error {
    lock.Lock()
    defer lock.Unlock()
    for i, cmd := range commandHistory {
        _, err := fmt.Fprintf(w, "%d %s\n", i+1, cmd)
        if err != nil {
            return err  // Return early if there's an error writing to the writer.
        }
    }
    return nil
}

