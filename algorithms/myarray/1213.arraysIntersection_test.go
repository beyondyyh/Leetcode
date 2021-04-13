package myarray

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Leetcode/algorithms/kit"
)

// run: go test -v 1213.*
func Test_arraysIntersection(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name: "x1",
			Input: [][]int{
				{1, 2, 3, 4, 5, 9},
				{2, 6, 7, 10},
				{1, 3, 4, 5, 8, 9, 10},
			},
			Expected: []int{},
		},
		{
			Name: "x2",
			Input: [][]int{
				{1, 2, 3, 4, 5, 9},
				{1, 2, 5, 7, 9},
				{1, 3, 4, 5, 8, 9, 10},
			},
			Expected: []int{1, 5, 9},
		},
	}

	assert := assert.New(t)
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			input := tt.Input.([][]int)
			assert.Equal(tt.Expected.([]int), arraysIntersection(input[0], input[1], input[2]))
		})
	}
}
