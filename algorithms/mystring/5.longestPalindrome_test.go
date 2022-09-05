package mystring

import (
	"testing"
)

// run: go test -v -run Test_longestPalindrome
func Test_longestPalindrome(t *testing.T) {
	cases := []struct {
		name, input string
		expected    []string
	}{
		{"x1", "babad", []string{"bab", "aba"}},
		{"x2", "cbbd", []string{"bb"}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output1 := longestPalindrome1(tt.input)
			if !stringContains(tt.expected, output1) {
				t.Errorf("longestPalindrome1(%s)=%s, expected is any of %s", tt.input, output1, tt.expected)
			}
			output2 := longestPalindrome2(tt.input)
			if !stringContains(tt.expected, output2) {
				t.Errorf("longestPalindrome2(%s)=%s, expected is any of %s", tt.input, output2, tt.expected)
			}
		})
	}
}

// stringContains returns true if a contains x
func stringContains(a []string, x string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return true
		}
	}
	return false
}
