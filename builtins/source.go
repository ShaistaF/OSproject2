package builtins

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// Source reads and executes commands from a file in the current shell environment.
func Source(filePath string) ([]byte, error) {
    // Open the file for reading.
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var output strings.Builder

    // Create a scanner to read lines from the file.
    scanner := bufio.NewScanner(file)

    // Iterate over each line in the file.
    for scanner.Scan() {
        // Get the command from the current line.
        line := scanner.Text()

        // Skip empty lines and comments.
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        // Execute the command and capture its output.
        cmd := exec.Command("sh", "-c", line)
        cmdOutput, err := cmd.CombinedOutput()
        if err != nil {
            return nil, fmt.Errorf("failed to execute command '%s': %v", line, err)
        }

        // Append the command output to the result.
        output.Write(cmdOutput)
    }

    // Check for any errors encountered during scanning.
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return []byte(output.String()), nil
}
