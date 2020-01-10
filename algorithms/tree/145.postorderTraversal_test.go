package tree

import (
	"reflect"
	"testing"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

type entry145 struct {
	name     string
	input    []int
	expected []int
}

// build_postorderTraversal_case returns test cases
func build_postorderTraversal_case() []entry145 {
	cases := []entry145{
		{
			name:     "x0",
			input:    []int{1, 2, 3},
			expected: []int{2, 3, 1},
		},
		{
			name:     "x1",
			input:    []int{1, NULL, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			name:     "x2",
			input:    []int{1, 2, 3, 4, 5, 6, NULL},
			expected: []int{4, 5, 2, 6, 3, 1},
		},
		{
			name:     "x3",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	return cases
}

// run: go test -v base.go 145.*
func Test_postorderTraversal(t *testing.T) {
	for _, tt := range build_postorderTraversal_case() {
		t.Run(tt.name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.input)
			if output := postorderTraversal(root); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("postorderTraversal(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}

func Test_Tree2Postorder(t *testing.T) {
	for _, tt := range build_postorderTraversal_case() {
		t.Run(tt.name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.input)
			if output := kit.Tree2Postorder(root); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("Tree2Postorder(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}
