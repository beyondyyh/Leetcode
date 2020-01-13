package mystring

import "testing"

// run: go test -v 3.*
func Test_lengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "x1",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "x2",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "x3",
			input:    "pwwkew",
			expected: 3,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := lengthOfLongestSubstring(tt.input); output != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%s)=%d, expected=%d", tt.input, output, tt.expected)
			}
		})
	}
}
