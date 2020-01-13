package mystring

import "testing"

// run: go test -v 32.*
func Test_longestValidParentheses1(t *testing.T) {
	cases := []struct {
		name, input string
		expected    int
	}{
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

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := longestValidParentheses1(tt.input); output != tt.expected {
				t.Errorf("longestValidParentheses1(%s)=%d, expected=%d", tt.input, output, tt.expected)
			}

			if output := longestValidParentheses2(tt.input); output != tt.expected {
				t.Errorf("longestValidParentheses2(%s)=%d, expected=%d", tt.input, output, tt.expected)
			}
		})
	}
}
