package tree

import (
	"reflect"
	"testing"
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

// run: go test -v treeNode.go 144.*
func Test_preorderTraversal(t *testing.T) {
	for _, tt := range build_preorderTraversal_case() {
		t.Run(tt.name, func(t *testing.T) {
			root := Ints2Tree(tt.input)
			if output := preorderTraversal(root); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("preorderTraversal(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}

func Test_preorderRecurse(t *testing.T) {
	for _, tt := range build_preorderTraversal_case() {
		t.Run(tt.name, func(t *testing.T) {
			root := Ints2Tree(tt.input)
			if output := preorderRecurse(root); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("preorderRecurse(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}
