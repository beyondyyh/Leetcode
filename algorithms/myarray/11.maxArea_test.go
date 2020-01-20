package myarray

import (
	"testing"

	"Leetcode/algorithms/kit"
)

func Test_maxArea(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Expected: 49,
		},
		{
			Name:     "x2",
			Input:    []int{1, 2, 3, 1},
			Expected: 3,
		},
		{
			Name:     "x3",
			Input:    []int{1, 1},
			Expected: 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			if output := maxArea(tt.Input.([]int)); output != tt.Expected.(int) {
				t.Errorf("maxArea(%v)=%d, expected=%d", tt.Input, output, tt.Expected.(int))
			}
		})
	}
}
