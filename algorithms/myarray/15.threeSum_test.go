package myarray

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Leetcode/algorithms/kit"
)

// run: go test -v 15.*
func Test_threeSum(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    []int{1, -1, -1, 0},
			Expected: [][]int{{-1, 0, 1}},
		},
		{
			Name:     "x2",
			Input:    []int{-1, 0, 1, 2, -1, -4},
			Expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	assert := assert.New(t)
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(tt.Expected.([][]int), threeSum(tt.Input.([]int)))
		})
	}
}
