package tree

import (
	"testing"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// go test -v base.go 101.*
func Test_isSymmetric(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "x1",
			input:    []int{1, 2, 2},
			expected: true,
		},
		{
			name:     "x2",
			input:    []int{1, 2, 2, 3, NULL, NULL, 3},
			expected: true,
		},
		{
			name:     "x3",
			input:    []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "x4",
			input:    []int{1, 2, 2, 3, 4, 4, 3},
			expected: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.input)
			if output := isSymmetric(root, root); output != tt.expected {
				t.Errorf("isSymmetric(%v)=%t, expected=%t", tt.input, output, tt.expected)
			}
		})
	}
}
