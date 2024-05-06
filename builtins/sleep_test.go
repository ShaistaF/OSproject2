package builtins

import (
	"bytes"
	"testing"
)

func TestSleep(t *testing.T) {
	w := &bytes.Buffer{}
	args := []string{"1"} 
	err := Sleep(w, args)
	if err != nil {
		t.Errorf("Sleep() returned an error: %v", err)
	}
	expected := "done\n"
	if got := w.String(); got != expected {
		t.Errorf("Sleep() = %q, want %q", got, expected)
	}
}