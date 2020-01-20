package tree

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v base.go 617.*
func Test_mergeTrees(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]int
		expected []interface{} // kit.BreadthFirstSearch returns []interface{}
	}{
		{
			name:     "x1",
			input:    [][]int{{1, 3, 2, 5}, {2, 1, 3, NULL, 4, NULL, 7}},
			expected: []interface{}{3, 4, 5, 5, 4, 7},
		},
		{
			name:     "x2",
			input:    [][]int{{1, 2, 3, 4, NULL, NULL, 5}, {1, 2, 3, NULL, 4, 5, NULL}},
			expected: []interface{}{2, 4, 6, 4, 4, 5, 5},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := kit.Ints2Tree(tt.input[0])
			t2 := kit.Ints2Tree(tt.input[1])
			output := mergeTrees(t1, t2)
			out2ints := kit.BreadthFirstSearch(*output)
			if !reflect.DeepEqual(out2ints, tt.expected) {
				t.Errorf("mergeTrees(%v)=%v, expected=%v", tt.input, out2ints, tt.expected)
			}
		})
	}
}
