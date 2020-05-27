package tree

import (
	"Leetcode/algorithms/kit"
	"testing"
)

// run: go test -v -run Test_isSubPath
func Test_isSubPath(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name: "x1",
			Input: map[string]interface{}{
				"head": []int{},
				"root": []int{1, 4, 4, null, 2, 2, null, 1, null, 6, 8, null, null, null, null, 1, 3},
			},
			Expected: true,
		},
		{
			Name: "x2",
			Input: map[string]interface{}{
				"head": []int{4, 2, 8},
				"root": []int{1, 4, 4, null, 2, 2, null, 1, null, 6, 8, null, null, null, null, 1, 3},
			},
			Expected: true,
		},
		{
			Name: "x3",
			Input: map[string]interface{}{
				"head": []int{1, 4, 2, 6},
				"root": []int{1, 4, 4, null, 2, 2, null, 1, null, 6, 8, null, null, null, null, 1, 3},
			},
			Expected: true,
		},
		{
			Name: "x4",
			Input: map[string]interface{}{
				"head": []int{1, 4, 2, 6, 8},
				"root": []int{1, 4, 4, null, 2, 2, null, 1, null, 6, 8, null, null, null, null, 1, 3},
			},
			Expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			input := c.Input.(map[string]interface{})
			head := kit.Ints2List(input["head"].([]int))
			root := kit.Ints2Tree(input["root"].([]int))
			if output := isSubPath(head, root); output != c.Expected.(bool) {
				t.Errorf("isSubPath()=%t, expected=%t", output, c.Expected)
			}
		})
	}
}
