package builtins

import (
    "bufio"
    "os"
    "os/exec"
    "strings"
)

// Source reads and executes commands from a file in the current shell environment.
func Source(filePath string) (string, error) {
    // Open the file for reading.
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
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
            return "", err
        }

        // Append the command output to the result.
        output.WriteString(string(cmdOutput))
        output.WriteString("\n") // Add a newline between command outputs
    }

    // Check for any errors encountered during scanning.
    if err := scanner.Err(); err != nil {
        return "", err
    }

    return output.String(), nil
}
