package tree

import (
	"reflect"
	"testing"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

type entry144 struct {
	name     string
	input    []int
	expected []int
}

func build_preorderTraversal_case() []entry144 {
	cases := []entry144{
		{
			name:     "x1",
			input:    []int{1, NULL, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "x2",
			input:    []int{1, 2, 3, 4, 5, 6, NULL},
			expected: []int{1, 2, 4, 5, 3, 6},
		},
		{
			name:     "x3",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	return cases
}

// run: go test -v base.go 144.*
func Test_preorderTraversal(t *testing.T) {
	for _, tt := range build_preorderTraversal_case() {
		t.Run(tt.name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.input)
			if output := preorderTraversal(root); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("preorderTraversal(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}

func Test_Tree2Preorder(t *testing.T) {
	for _, tt := range build_preorderTraversal_case() {
		t.Run(tt.name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.input)
			if output := kit.Tree2Preorder(root); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("Tree2Preorder(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}
