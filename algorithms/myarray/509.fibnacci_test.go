package myarray

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// run: go test -v 509.*
func Test_fibnacci(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    2,
			Expected: 1,
		},
		{
			Name:     "x2",
			Input:    3,
			Expected: 2,
		},
		{
			Name:     "x3",
			Input:    4,
			Expected: 3,
		},
		{
			Name:     "x4",
			Input:    10,
			Expected: 55,
		},
	}

	assert := assert.New(t)
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(tt.Expected.(int), fibnacci(tt.Input.(int)))
		})
	}
}
