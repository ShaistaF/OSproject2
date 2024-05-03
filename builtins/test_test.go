package builtins_test

import (
	"bytes"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestTest_NoArgs(t *testing.T) {
	var buf bytes.Buffer

	err := builtins.Test(&buf)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	expectedErr := "no test expression provided"
	if err.Error() != expectedErr {
		t.Errorf("Expected error: %s, got: %v", expectedErr, err)
	}
}

func TestTest_WithArgs(t *testing.T) {
	var buf bytes.Buffer

	expression := "1 -eq 1"
	err := builtins.Test(&buf, expression)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Test expression: 1 -eq 1\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, buf.String())
	}
}

func TestTest_MultipleArgs(t *testing.T) {
	var buf bytes.Buffer

	expression := "1 -eq 1 && 2 -eq 2"
	err := builtins.Test(&buf, expression)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Test expression: 1 -eq 1 && 2 -eq 2\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, buf.String())
	}
}

func TestTest_ErrorOnEvaluation(t *testing.T) {
	var buf bytes.Buffer

	expression := "1 -eq 2"
	err := builtins.Test(&buf, expression)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Test expression: 1 -eq 2\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, buf.String())
	}
}

func TestTest_ErrorOnInvalidExpression(t *testing.T) {
	var buf bytes.Buffer

	expression := "1 eq 2"
	err := builtins.Test(&buf, expression)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	expectedErr := "invalid expression: 1 eq 2"
	if err.Error() != expectedErr {
		t.Errorf("Expected error: %s, got: %v", expectedErr, err)
	}
}

func TestTest_GreaterThan(t *testing.T) {
	var buf bytes.Buffer

	expression := "2 -gt 1"
	err := builtins.Test(&buf, expression)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Test expression: 2 -gt 1\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, buf.String())
	}
}

func TestTest_LessThan(t *testing.T) {
	var buf bytes.Buffer

	expression := "1 -lt 2"
	err := builtins.Test(&buf, expression)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Test expression: 1 -lt 2\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, buf.String())
	}
}

func TestTest_NotEqual(t *testing.T) {
	var buf bytes.Buffer

	expression := "1 -ne 2"
	err := builtins.Test(&buf, expression)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Test expression: 1 -ne 2\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, buf.String())
	}
}
