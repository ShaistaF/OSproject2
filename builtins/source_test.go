package builtins

import (
    "bytes"
    "io/ioutil"
    "os"
    "testing"
)

func TestSource(t *testing.T) {
    // Create a temporary script file
    tmpFile, err := ioutil.TempFile("", "example.*.sh")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpFile.Name())

    content := "echo Hello World\npwd"
    if _, err = tmpFile.WriteString(content); err != nil {
        t.Fatal(err)
    }
    tmpFile.Close()

    w := new(bytes.Buffer)
    Source(w, tmpFile.Name())
    expectedOutput := "echo Hello World\npwd\n"
    if got := w.String(); got != expectedOutput {
        t.Errorf("Expected %q, got %q", expectedOutput, got)
    }
}
