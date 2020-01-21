package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Leetcode/algorithms/kit"
)

// go test -v base.go 102.*
func Test_levelOrder(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:  "x1",
			Input: []int{3, 9, 20, NULL, NULL, 15, 7},
			Expected: [][]int{
				[]int{3},
				[]int{9, 20},
				[]int{15, 7},
			},
		},
		{
			Name:  "x2",
			Input: []int{3, 9, 20, 10, 11, 15, 7},
			Expected: [][]int{
				[]int{3},
				[]int{9, 20},
				[]int{10, 11, 15, 7},
			},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := kit.Ints2Tree(c.Input.([]int))
			assert.Equal(c.Expected, levelOrder1(root), "Input:%v", c.Input)
			assert.Equal(c.Expected, levelOrder2(root), "Input:%v", c.Input)
		})
	}
}
