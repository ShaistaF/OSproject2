package builtins

import (
    "bytes"
    "testing"
)

func TestTestCommand(t *testing.T) {
    tests := []struct {
        args     []string
        expected string
    }{
        {[]string{"5", "-eq", "5"}, "true\n"},
        {[]string{"5", "-ne", "4"}, "true\n"},
        {[]string{"3", "-gt", "4"}, "false\n"},
        {[]string{"4", "-ge", "4"}, "true\n"},
        {[]string{"5", "-lt", "6"}, "true\n"},
        {[]string{"7", "-le", "7"}, "true\n"},
        {[]string{"5", "-le", "3"}, "false\n"},
        {[]string{"a", "-eq", "b"}, "test: both operands must be integers for comparison\n"},
        {[]string{"5"}, "test: expected three arguments for comparisons\n"},
    }

    for _, tt := range tests {
        w := new(bytes.Buffer)
        Test(w, tt.args...)
        if got := w.String(); got != tt.expected {
            t.Errorf("Test(%q) = %q, want %q", tt.args, got, tt.expected)
        }
    }
}
