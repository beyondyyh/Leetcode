package myarray

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Leetcode/algorithms/kit"
)

// run: go test -v 53.*
func Test_maxSubArray(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			Expected: 6,
		},
		{
			Name:     "x2",
			Input:    []int{-1, -2},
			Expected: -1,
		},
	}

	assert := assert.New(t)
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(tt.Expected.(int), maxSubArray(tt.Input.([]int)))
		})
	}
}
