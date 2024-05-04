package builtins

import (
    "bytes"
    "testing"
)

func TestEcho(t *testing.T) {
    cases := []struct {
        name     string
        args     []string
        expected string
    }{
        {"no arguments", []string{}, "\n"},
        {"single word", []string{"Hello"}, "Hello\n"},
        {"multiple words", []string{"Hello", "world!"}, "Hello world!\n"},
        {"handles spaces", []string{"Hello", "", "world!"}, "Hello  world!\n"},
        {"numbers and symbols", []string{"123", "&*^%"}, "123 &*^%\n"},
        {"long string", []string{"This", "is", "a", "very", "long", "sentence", "with", "many", "words."},
            "This is a very long sentence with many words.\n"},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            w := &bytes.Buffer{}
            err := Echo(w, tc.args...)
            if err != nil {
                t.Fatalf("Echo() returned an error: %v", err)
            }
            if got := w.String(); got != tc.expected {
                t.Errorf("Echo() = %q, want %q", got, tc.expected)
            }
        })
    }
}
