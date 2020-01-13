package mystring

import "testing"

// run: go test -v 14.*
func Test_longestCommonPrefix(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "x1",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "x2",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := longestCommonPrefix(tt.input); output != tt.expected {
				t.Errorf("longestCommonPrefix(%v)=%s, expected=%s", tt.input, output, tt.expected)
			}
		})
	}
}
