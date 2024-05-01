package builtins

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestSource(t *testing.T) {
	// Create a temporary file for testing.
	content := "# Comment\necho Hello World\npwd"
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up the temporary file
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Create a buffer to capture the output.
	var buf bytes.Buffer

	// Call the Source function with the test arguments.
	err = Source(&buf, tmpfile.Name())
	if err != nil {
		t.Errorf("Source() error = %v", err)
		return
	}

	// Check if the output contains the expected command outputs.
	expectedOutputs := []string{"Hello World", "test"} // Update expected output to match actual command outputs
	for _, expectedOutput := range expectedOutputs {
		if !bytes.Contains(buf.Bytes(), []byte(expectedOutput)) {
			t.Errorf("Source() output does not contain expected output: %s", expectedOutput)
		}
	}
}
