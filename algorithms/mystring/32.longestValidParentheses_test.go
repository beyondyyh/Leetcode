package main

import "testing"

type entry32 struct {
	name     string
	input    string
	expected int
}

// build test case
func build_longsestValidParentheses_case() []entry32 {
	cases := []entry32{
		{
			name:     "x1",
			input:    "(()",
			expected: 2,
		},
		{
			name:     "x2",
			input:    ")()())",
			expected: 4,
		},
		{
			name:     "x3",
			input:    "((())",
			expected: 4,
		},
	}

	return cases
}

// run: go test -v 32.*
func Test_longestValidParentheses1(t *testing.T) {
	tests := build_longsestValidParentheses_case()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if output := longestValidParentheses1(tt.input); output != tt.expected {
				t.Errorf("longestValidParentheses1(%s)=%d, expected=%d", tt.input, output, tt.expected)
			}
		})
	}
}

func Test_longestValidParentheses2(t *testing.T) {
	tests := build_longsestValidParentheses_case()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if output := longestValidParentheses2(tt.input); output != tt.expected {
				t.Errorf("longestValidParentheses2(%s)=%d, expected=%d", tt.input, output, tt.expected)
			}
		})
	}
}
