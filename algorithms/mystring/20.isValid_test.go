package mystring

import "testing"

// run: go test -v 20.*
func Test_isValid(t *testing.T) {
	cases := []struct {
		name, input string
		expected    bool
	}{
		{"x1", "()", true},
		{"x2", "()[]{}", true},
		{"x3", "(]", false},
		{"x4", "([)]", false},
		{"x5", "{[]}", true},
		{"x6", "()()((([]{})))", true},
		{"x7", "()()())", false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := isValid1(tt.input); output != tt.expected {
				t.Errorf("isValid1(%s)=%t, expected=%t", tt.input, output, tt.expected)
			}
			if output := isValid2(tt.input); output != tt.expected {
				t.Errorf("isValid2(%s)=%t, expected=%t", tt.input, output, tt.expected)
			}
			if output := isValid3(tt.input); output != tt.expected {
				t.Errorf("isValid3(%s)=%t, expected=%t", tt.input, output, tt.expected)
			}
		})
	}
}
