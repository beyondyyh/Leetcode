package tree

import (
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v -run Test_isBalanced
func Test_isBalanced(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    []int{3, 9, 20, null, null, 15, 7},
			Expected: true,
		},
		{
			Name:     "x2",
			Input:    []int{1, 2, 2, 3, 3, null, null, 4, 4},
			Expected: false,
		},
		{
			Name:     "x3",
			Input:    []int{},
			Expected: true,
		},
		// add more test cases
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.Input.([]int))
			if output := isBalanced(root); output != tt.Expected.(bool) {
				t.Errorf("isBalanced(%v)=%t, expected=%v", tt.Input, output, tt.Expected)
			}
		})
	}
}
