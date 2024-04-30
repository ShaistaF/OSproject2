func TestTest(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr bool
	}{
		{"Equal Test", []string{"hello", "=", "hello"}, "true\n", false},
		{"Not Equal Test", []string{"hello", "!=", "world"}, "true\n", false},
		{"Greater Than Test", []string{"5", ">", "3"}, "true\n", false},
		{"Less Than Test", []string{"3", "<", "5"}, "true\n", false},
		{"Greater Than Or Equal To Test", []string{"5", ">=", "5"}, "true\n", false},
		{"Less Than Or Equal To Test", []string{"5", "<=", "5"}, "true\n", false},
		{"Invalid Operator Test", []string{"hello", "+", "world"}, "", true},
		{"Missing Args Test", []string{"hello", "="}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture the output.
			var buf bytes.Buffer

			// Call the Test function with the test arguments and the buffer as the writer.
			err := Test(&buf, tt.args...)

			// Check if there was an error.
			if (err != nil) != tt.expectedErr {
				t.Errorf("Test() error = %v, expected error = %v", err, tt.expectedErr)
				return
			}

			// Check if the output matches the expected output.
			if buf.String() != tt.expectedOut {
				t.Errorf("Test() = %v, want %v", buf.String(), tt.expectedOut)
			}
		})
	}
}
