package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/ShaistaF/OSproject2/builtins"
)

func main() {
	exit := make(chan struct{}, 2)
	runLoop(os.Stdin, os.Stdout, os.Stderr, exit)
}

func runLoop(r io.Reader, w, errW io.Writer, exit chan struct{}) {
	var (
		input    string
		err      error
		readLoop = bufio.NewReader(r)
	)
	for {
		select {
		case <-exit:
			_, _ = fmt.Fprintln(w, "exiting gracefully...")
			return
		default:
			if err := printPrompt(w); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if input, err = readLoop.ReadString('\n'); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if err = handleInput(w, input, exit); err != nil {
				_, _ = fmt.Fprintln(errW, err)
			}
		}
	}
}

func printPrompt(w io.Writer) error {
	// Get current user.
	// Don't prematurely memoize this because it might change due to `su`?
	u, err := user.Current()
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%v [%v] $ ", wd, u.Username)

	return err
}

func handleInput(w io.Writer, input string, exit chan<- struct{}) error {
	input = strings.TrimSpace(input)

	args := strings.Split(input, " ")
	name, args := args[0], args[1:]

	switch name {
	case "cd":
		return builtins.ChangeDirectory(args...)
	case "env":
		return builtins.EnvironmentVariables(w, args...) //not working?
	case "echo":
		return builtins.Echo(w, args...)
	case "whoami":
    		return builtins.Whoami(w)
	case "clear":
    		return builtins.Clear(w)
	case "date":
			return builtins.Date(w)
	case "history":
    		return builtins.History(w)
	case "pwd":
			return builtins.Pwd(w)
	case "sleep":
			return builtins.Sleep(w, args)
	case "hostname":
			return builtins.Hostname(w)
	case "version":
			return builtins.Version(w)
	case "help":
			return builtins.Help(w) //edit, update
		
		

	case "exit":
		exit <- struct{}{}
		return nil
	}

	return executeCommand(name, args...)
}

func executeCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

