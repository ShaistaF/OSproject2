package builtins

import (
	"fmt"
	"io"
)

func Echo(w io.Writer, args ...string) error {
	for i, arg := range args {
		if i > 0 {
			_, err := fmt.Fprint(w, " ")
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprint(w, arg)
		if err != nil {
			return err
		}
	}

	
	_, err := fmt.Fprintln(w)
	return err
}
