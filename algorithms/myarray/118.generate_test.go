package myarray

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// go test -v 118.*
func Test_generate(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:  "x1",
			Input: 1,
			Expected: [][]int{
				[]int{1},
			},
		},
		{
			Name:  "x2",
			Input: 5,
			Expected: [][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			},
		},
	}

	assert := assert.New(t)
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			output := generate(tt.Input.(int))
			assert.Equal(tt.Expected, output, "generate(%d)=%v, expected=%v", tt.Input, output, tt.Expected)
		})
	}
}
