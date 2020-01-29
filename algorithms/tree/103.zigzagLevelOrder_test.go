package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Leetcode/algorithms/kit"
)

// run: go test -v base.go 103.*
func Test_zigzagLevelOrder(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    []int{3, 9, 20, null, null, 15, 7},
			Expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			Name:     "x2",
			Input:    []int{3, 9, 20, 15, 7},
			Expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := kit.Ints2Tree(c.Input.([]int))
			output := zigzagLevelOrder2(root)
			assert.Equal(c.Expected.([][]int), output, "expected=%v, output=%v", c.Expected, output)
		})
	}
}
