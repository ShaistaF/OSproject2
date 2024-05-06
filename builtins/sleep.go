package builtins

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

func Sleep(w io.Writer, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: sleep <seconds>")
	}
	seconds, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid number of seconds: %s", args[0])
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	_, err = fmt.Fprintln(w, "done")
	return err
}

