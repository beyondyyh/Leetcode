package linkedList

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v base.go 21.*
func Test_mergeTwoSortedLists(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "x1",
			input:    [][]int{{1, 2, 4}, {1, 3, 4}},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "x2",
			input:    [][]int{{1, 2, 4}, {}},
			expected: []int{1, 2, 4},
		},
		{
			name:     "x3",
			input:    [][]int{{1, 3, 5}, {2, 4, 6, 7, 8}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			l1 := kit.Ints2List(tt.input[0])
			l2 := kit.Ints2List(tt.input[1])
			output := mergeTwoSortedLists(l1, l2)
			out2ints := kit.List2Ints(output)
			if !reflect.DeepEqual(out2ints, tt.expected) {
				t.Errorf("mergeTwoSortedLists(%v, %v)=%v, expected=%v", tt.input[0], tt.input[1], out2ints, tt.expected)
			}
		})
	}
}
