package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Command executes the specified command with the given arguments, ignoring shell functions.
func Command(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("no command specified")
	}

	opts, args := parseOptions(args)
	if len(args) == 0 {
		return fmt.Errorf("no command specified after options")
	}

	commandName, commandArgs := args[0], args[1:]

	// Handle options
	switch {
	case strings.Contains(opts, "V"):
		if path, err := exec.LookPath(commandName); err == nil {
			fmt.Printf("%s is located at %s\n", commandName, path)
		} else {
			fmt.Println("command not found")
			return err
		}
		return nil
	case strings.Contains(opts, "v"):
		fmt.Println(commandName)
		return nil
	case strings.Contains(opts, "p"):
		os.Setenv("PATH", "/bin:/usr/bin:/usr/local/bin")
	}

	// Execute the command
	cmd := exec.Command(commandName, commandArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("command failed with status %d", exiterr.ExitCode())
		}
		return fmt.Errorf("command execution failed: %v", err)
	}
	return nil
}

func parseOptions(args []string) (string, []string) {
	opts := ""
	index := 0
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 && isOption(arg[1:]) {
			opts += arg[1:]
			index++
		} else {
			break
		}
	}
	return opts, args[index:]
}

func isOption(s string) bool {
	for _, r := range s {
		if r != 'p' && r != 'V' && r != 'v' {
			return false
		}
	}
	return true
}

