package builtins

import (
    "fmt"
    "io"
    "os/user"
)

// Whoami prints the current system user's username to the provided io.Writer.
func Whoami(w io.Writer) error {
    usr, err := user.Current()
    if err != nil {
        return fmt.Errorf("whoami: unable to fetch user: %v", err)
    }
    _, err = fmt.Fprintln(w, usr.Username)
    return err
}
