package tree

import (
	"testing"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// run: go test -v base.go 104.*
func Test_maxDepth(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "x1",
			input:    []int{3, 9, 20, NULL, NULL, 15, 7},
			expected: 3,
		},
		{
			name:     "x2",
			input:    []int{3, 9, 20},
			expected: 2,
		},
		{
			name:     "x3",
			input:    []int{3, 9, 20, 15, 7},
			expected: 3,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.input)
			if output := maxDepth(root); output != tt.expected {
				t.Errorf("maxDepth(%v)=%d, expected=%d", tt.input, output, tt.expected)
			}
		})
	}
}
