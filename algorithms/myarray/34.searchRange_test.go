package myarray

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

type entry34input struct {
	nums   []int
	target int
}

// run: go test -v 34.*
func Test_searchRange(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name: "x1",
			Input: entry34input{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			Expected: []int{3, 4},
		},
		{
			Name: "x2",
			Input: entry34input{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			Expected: []int{-1, -1},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			input := tt.Input.(entry34input)
			if output := searchRange(input.nums, input.target); !reflect.DeepEqual(output, tt.Expected.([]int)) {
				t.Errorf("searchRange(%v, %d)=%d, expected=%d", input.nums, input.target, output, tt.Expected)
			}
		})
	}
}
