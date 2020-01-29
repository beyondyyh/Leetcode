package mybit

import (
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v 136.*
func Test_singleNumber(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    []int{2, 2, 1},
			Expected: 1,
		},
		{
			Name:     "x2",
			Input:    []int{4, 1, 2, 1, 2},
			Expected: 4,
		},
		{
			Name:     "x3",
			Input:    []int{1, 2, 3, 3, 2, 1, 1},
			Expected: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if output := singleNumber(c.Input.([]int)); c.Expected.(int) != output {
				t.Errorf("singleNumber(%v)=%d, expected=%d", c.Input, output, c.Expected)
			}
		})
	}
}
