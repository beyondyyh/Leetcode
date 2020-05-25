package tree

import (
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v -run Test_lowestCommonAncestor
func Test_lowestCommonAncestor(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name: "x1",
			Input: map[string]interface{}{
				"root": []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
				"p":    5,
				"q":    1,
			},
			Expected: 3,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			input := tt.Input.(map[string]interface{})
			root := kit.Ints2Tree(input["root"].([]int))
			p := kit.FindTargetNode(root, input["p"].(int))
			q := kit.FindTargetNode(root, input["q"].(int))
			expected := kit.FindTargetNode(root, tt.Expected.(int))
			if output := lowestCommonAncestor(root, p, q); output != expected {
				t.Errorf("lowestCommonAncestor(%v, %d, %d)=%d, expected=%d", input["root"], p.Val, q.Val, output.Val, expected.Val)
			}
		})
	}
}
