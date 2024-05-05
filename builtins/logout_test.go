package builtins

import (
	"testing"
)

func TestLogout(t *testing.T) {
	tests := []struct {
		args    []string
		wantErr bool
		errMsg  string
	}{
		{[]string{"0"}, false, ""},
		{[]string{"1"}, false, ""},
		{[]string{"-1"}, false, ""},
		{[]string{"a"}, true, "invalid exit status: a"},
		{[]string{}, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.args[0], func(t *testing.T) {
			err := Logout(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logout(%v) error = %v, wantErr %v", tt.args, err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("Logout() error = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
