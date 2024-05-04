package builtins

import (
    "fmt"
    "io"
    "sync"
)


var commandHistory []string
var lock sync.Mutex


func AddToHistory(command string) {
    lock.Lock()
    defer lock.Unlock()
    commandHistory = append(commandHistory, command)
}


func History(w io.Writer) error {
    lock.Lock()
    defer lock.Unlock()
    for i, cmd := range commandHistory {
        _, err := fmt.Fprintf(w, "%d %s\n", i+1, cmd)
        if err != nil {
            return err  
        }
    }
    return nil
}

