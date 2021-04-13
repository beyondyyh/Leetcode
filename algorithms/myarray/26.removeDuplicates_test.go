package myarray

import (
	"testing"
)

// run: go test -v 26.*
func Test_removeDuplicates(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "case 1",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "case 2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// if got := removeDuplicates1(c.input); got != c.expected {
			// 	t.Errorf("removeDuplicates1(%v)=%v, expected=%v", c.input, got, c.expected)
			// }
			if got := removeDuplicates2(c.input); got != c.expected {
				t.Errorf("removeDuplicates2(%v)=%v, expected=%v", c.input, got, c.expected)
			}
		})
	}
}
