package mystring

import "testing"

// run: go test -v 1108.*
func Test_defangIPaddr(t *testing.T) {
	cases := []struct {
		name, input, expected string
	}{
		{"x1", "1.1.1.1", "1[.]1[.]1[.]1"},
		{"x2", "255.100.50.0", "255[.]100[.]50[.]0"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := defangIPaddr(tt.input); output != tt.expected {
				t.Errorf("defangIPaddr(%s)=%s, expected=%s", tt.input, output, tt.expected)
			}
		})
	}
}
