package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)


func Export(w io.Writer, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: export name1=value1 [name2=value2 ...]")
	}

	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid argument: %s", arg)
		}
		key, value := parts[0], parts[1]
		os.Setenv(key, value)
		_, err := fmt.Fprintf(w, "Exported variable: %s=%s\n", key, value)
		if err != nil {
			return err
		}
	}

	return nil
}
