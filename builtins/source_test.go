package builtins

import (
    "io/ioutil"
    "os"
    "testing"
)

// TestSource tests the Source function.
func TestSource(t *testing.T) {
    // Create a temporary file with commands
    file, err := ioutil.TempFile("", "testfile.txt")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(file.Name()) // Clean up

    // Write commands to the file
    commands := []string{"echo Hello", "echo World"}
    for _, cmd := range commands {
        _, err := file.WriteString(cmd + "\n")
        if err != nil {
            t.Fatal(err)
        }
    }
    file.Close()

    // Call the Source function
    output, err := Source(file.Name())

    // Check for error
    if err != nil {
        t.Fatalf("Source error: %v", err)
    }

    // Define the expected output
    expected := []byte("Hello\nWorld\n")

    // Compare the actual output with the expected output
    if string(output) != string(expected) {
        t.Errorf("Source() output does not match expected output: got %q, want %q", output, expected)
    }
}
