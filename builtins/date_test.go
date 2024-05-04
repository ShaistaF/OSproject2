// date_test.go
package builtins

import (
    "bytes"
    "testing"
    "time"
)

func TestDate(t *testing.T) {
    var buf bytes.Buffer
    expectedDate := time.Now().Format("2006-01-02") // Ensure the test checks for the current date in ISO 8601 format

    if err := Date(&buf); err != nil {
        t.Fatalf("Date returned an error: %v", err)
    }
    if got := buf.String(); got != expectedDate {
        t.Errorf("Expected date to be %v, got %v", expectedDate, got)
    }
}
