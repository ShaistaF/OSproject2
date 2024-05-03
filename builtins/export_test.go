package builtins_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestExport_Print(t *testing.T) {
	var buf bytes.Buffer

	// Test printing environment variables
	err := builtins.Export(&buf)
	if err != nil {
		t.Errorf("Export failed: %v", err)
	}

	// Split buffer by lines to count environment variables
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != len(os.Environ()) {
		t.Errorf("Expected %d environment variables, got %d", len(os.Environ()), len(lines))
	}
}

func TestExport_Set(t *testing.T) {
	// Test setting environment variables
	err := builtins.Export(nil, "key1=value1", "key2=value2")
	if err != nil {
		t.Errorf("Export failed: %v", err)
	}

	// Check if environment variables are set
	val1 := os.Getenv("key1")
	if val1 != "value1" {
		t.Errorf("Environment variable key1 not set correctly: got %s, want value1", val1)
	}

	val2 := os.Getenv("key2")
	if val2 != "value2" {
		t.Errorf("Environment variable key2 not set correctly: got %s, want value2", val2)
	}
}

func TestExport_InvalidArgumentFormat(t *testing.T) {
	// Test invalid argument format
	err := builtins.Export(nil, "invalid_format")
	if err == nil {
		t.Error("Export should return error for invalid argument format")
	}
}
