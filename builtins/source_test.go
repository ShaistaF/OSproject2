package builtins_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestSource_Success(t *testing.T) {
	// Create a temporary file with shell commands
	file, err := ioutil.TempFile("", "test_source_*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write shell commands to the file
	commands := []string{
		"echo 'Hello, World!'",
		"ls -l",
		"echo 'Goodbye!'",
	}
	for _, cmd := range commands {
		if _, err := file.WriteString(cmd + "\n"); err != nil {
			t.Fatal(err)
		}
	}

	var buf bytes.Buffer

	// Execute the source command on the temporary file
	err = builtins.Source(&buf, file.Name())
	if err != nil {
		t.Errorf("Source failed: %v", err)
	}

	// Check if the output contains the executed commands
	for _, cmd := range commands {
		if !bytes.Contains(buf.Bytes(), []byte(cmd)) {
			t.Errorf("Output does not contain command: %s", cmd)
		}
	}
}

func TestSource_FileNotFound(t *testing.T) {
	var buf bytes.Buffer

	// Execute the source command on a non-existent file
	err := builtins.Source(&buf, "non_existent_file.sh")
	if err == nil {
		t.Error("Expected error, got nil")
	}
	expectedErr := "open non_existent_file.sh: no such file or directory"
	if !errors.Is(err, os.ErrNotExist) || err.Error() != expectedErr {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}

func TestSource_EmptyFile(t *testing.T) {
	// Create an empty temporary file
	file, err := ioutil.TempFile("", "empty_source_*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	var buf bytes.Buffer

	// Execute the source command on the empty file
	err = builtins.Source(&buf, file.Name())
	if err != nil {
		t.Errorf("Source failed: %v", err)
	}

	// Check if the output is empty
	if buf.Len() > 0 {
		t.Errorf("Expected empty output for empty file, got: %s", buf.String())
	}
}

func TestSource_InvalidCommands(t *testing.T) {
	// Create a temporary file with invalid shell commands
	file, err := ioutil.TempFile("", "invalid_source_*.sh")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write invalid shell commands to the file
	invalidCommands := []string{
		"invalid command",
		"echo 'Hello, World!' && invalid_command",
	}
	for _, cmd := range invalidCommands {
		if _, err := file.WriteString(cmd + "\n"); err != nil {
			t.Fatal(err)
		}
	}

	var buf bytes.Buffer

	// Execute the source command on the temporary file
	err = builtins.Source(&buf, file.Name())
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Check if the error indicates failed command execution
	expectedErr := "failed to execute command"
	if !bytes.Contains([]byte(err.Error()), []byte(expectedErr)) {
		t.Errorf("Expected error containing '%s', got: %v", expectedErr, err)
	}
}
