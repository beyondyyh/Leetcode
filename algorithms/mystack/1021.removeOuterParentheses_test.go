package mystack

import (
	"testing"
)

// run: go test -v base.go 1021.*
func Test_removeOuterParentheses(t *testing.T) {
	cases := []struct {
		name, input, expected string
	}{
		{
			name:     "x1",
			input:    "(()())(())",
			expected: "()()()",
		},
		{
			name:     "x2",
			input:    "(()())(())(()(()))",
			expected: "()()()()(())",
		},
		{
			name:     "x3",
			input:    "()()",
			expected: "",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := removeOuterParentheses(tt.input); output != tt.expected {
				t.Errorf("removeOuterParentheses(\"%s\")=%s, expected=%s", tt.input, output, tt.expected)
			}
		})
	}
}
