package builtins_test

import (
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestAlias(t *testing.T) {
	// Set an alias
	builtins.SetAlias("ll", "ls -l")

	// Get and check the alias
	val, ok := builtins.GetAlias("ll")
	if !ok || val != "ls -l" {
		t.Errorf("SetAlias failed: got %s, want %s", val, "ls -l")
	}

	// Test setting an existing alias
	err := builtins.SetAlias("ll", "ls -l -h")
	if err != builtins.ErrAliasExists {
		t.Errorf("SetAlias should return ErrAliasExists: got %v", err)
	}

	// Unset the alias
	builtins.UnsetAlias("ll")

	// Get and check the unset alias
	_, ok = builtins.GetAlias("ll")
	if ok {
		t.Errorf("UnsetAlias failed: alias still exists")
	}

	// Test unsetting a non-existing alias
	builtins.UnsetAlias("nonexistent")
	_, ok = builtins.GetAlias("nonexistent")
	if ok {
		t.Errorf("UnsetAlias failed: alias still exists")
	}

	// Test getting an unset alias
	_, err = builtins.GetAlias("unset")
	if err != builtins.ErrAliasNotSet {
		t.Errorf("GetAlias should return ErrAliasNotSet: got %v", err)
	}

	// Set another alias
	builtins.SetAlias("grep", "grep --color=auto")

	// Test expanding alias
	cmd := "grep pattern file.txt"
	expandedCmd := builtins.ExpandAlias(cmd)
	expectedCmd := "grep --color=auto pattern file.txt"
	if expandedCmd != expectedCmd {
		t.Errorf("ExpandAlias failed: got %s, want %s", expandedCmd, expectedCmd)
	}

	// Test printing aliases
	builtins.PrintAliases()
}
