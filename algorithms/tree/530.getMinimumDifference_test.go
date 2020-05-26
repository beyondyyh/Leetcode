package tree

import (
	"Leetcode/algorithms/kit"
	"testing"
)

// run: go test -v -run Test_getMinimumDifference
func Test_getMinimumDifference(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    []int{1, null, 3, 2},
			Expected: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := kit.Ints2Tree(c.Input.([]int))
			if output := getMinimumDifference(root); output != c.Expected.(int) {
				t.Errorf("getMinimumDifference(%v)=%d, expected=%d\n", c.Input, output, c.Expected)
			}
		})
	}
}
