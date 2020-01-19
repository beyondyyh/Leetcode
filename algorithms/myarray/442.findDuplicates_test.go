package myarray

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// run: go test -v 442.*
func Test_findDuplicates(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    []int{4, 3, 2, 7, 8, 2, 3, 1},
			Expected: []int{2, 3},
		},
		{
			Name:     "x2",
			Input:    []int{1, 2, 3, 4, 5},
			Expected: []int{},
		},
	}

	assert := assert.New(t)
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(tt.Expected, findDuplicates(tt.Input.([]int)))
		})
	}
}
