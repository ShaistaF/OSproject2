package builtins

import (
	"errors"
	"fmt"
	"strings"
)

// Alias creates or displays aliases.
type AliasMap map[string]string

var aliases AliasMap

// SetAlias sets an alias.
func SetAlias(name, value string) {
	if aliases == nil {
		aliases = make(AliasMap)
	}
	aliases[name] = value
}

// GetAlias retrieves the value of an alias.
func GetAlias(name string) (string, bool) {
	val, ok := aliases[name]
	return val, ok
}

// UnsetAlias removes an alias.
func UnsetAlias(name string) {
	delete(aliases, name)
}

// PrintAliases prints all aliases to standard output.
func PrintAliases() {
	for name, value := range aliases {
		fmt.Printf("%s=%s\n", name, value)
	}
}

// ExpandAlias expands the alias in a command.
func ExpandAlias(cmd string) string {
	parts := strings.Fields(cmd)
	for i, part := range parts {
		if val, ok := aliases[part]; ok {
			parts[i] = val
		}
	}
	return strings.Join(parts, " ")
}

// ErrAliasNotFound is returned when alias is not found.
var ErrAliasNotFound = errors.New("alias not found")

// ErrAliasExists is returned when trying to set an alias that already exists.
var ErrAliasExists = errors.New("alias already exists")

// ErrAliasNotSet is returned when trying to get an unset alias.
var ErrAliasNotSet = errors.New("alias not set")
