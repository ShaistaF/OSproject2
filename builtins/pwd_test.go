package builtins

import (
	"bytes"
	"os"   
	"testing"
)

func TestPwd(t *testing.T) {
	w := &bytes.Buffer{}
	err := Pwd(w)
	if err != nil {
		t.Fatalf("Pwd() returned an error: %v", err)
	}
	expected, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getting current directory failed: %v", err)
	}
	if got := w.String(); got != expected+"\n" {
		t.Errorf("Pwd() = %q, want %q", got, expected)
	}
}