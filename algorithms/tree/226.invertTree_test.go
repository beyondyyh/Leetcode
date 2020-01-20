package tree

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v base.go 226.*
func Test_invertTree(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []interface{} // kit.BreadthFirstSearch returns []interface{}
	}{
		{
			name:     "x1",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []interface{}{1, 3, 2, 7, 6, 5, 4},
		},
		{
			name:     "x2",
			input:    []int{4, 2, 7, 1, 3, 6, 9},
			expected: []interface{}{4, 7, 2, 9, 6, 3, 1},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output := invertTree(kit.Ints2Tree(tt.input))
			out2ints := kit.BreadthFirstSearch(*output)
			if !reflect.DeepEqual(out2ints, tt.expected) {
				t.Errorf("invertTree(%v)=%v, expected=%v", tt.input, out2ints, tt.expected)
			}
		})
	}
}
