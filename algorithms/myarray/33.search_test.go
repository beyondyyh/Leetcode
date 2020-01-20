package myarray

import (
	"testing"

	"Leetcode/algorithms/kit"
)

type entry33input struct {
	nums   []int
	target int
}

// run: go test -v 33.*
func Test_search(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name: "x1",
			Input: entry33input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			Expected: 4,
		},
		{
			Name: "x2",
			Input: entry33input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			Expected: -1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			input := tt.Input.(entry33input)
			if output := search(input.nums, input.target); output != tt.Expected.(int) {
				t.Errorf("search(%v)=%d, expected=%d", input, output, tt.Expected)
			}
		})
	}
}
