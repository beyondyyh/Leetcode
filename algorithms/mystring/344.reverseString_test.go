package mystring

import "testing"

// run: go test -v 344.*
func Test_reverseString(t *testing.T) {
	cases := []struct {
		name            string
		input, expected []byte
	}{
		{"x1", []byte("hello"), []byte("olleh")},
		{"x2", []byte("Hannah"), []byte("hannaH")},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := reverseString(tt.input); output != string(tt.expected) {
				t.Errorf("reverseString(%s)=%s, expected=%s", string(tt.input), string(output), string(tt.expected))
			}
		})
	}
}
