package linkedList

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v base.go 206.*
func Test_reverseList(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		// {
		// 	name:     "x1",
		// 	input:    []int{},
		// 	expected: []int{},
		// },
		{
			name:     "x2",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			head := kit.Ints2List(tt.input)
			output := reverseListV2(head)
			out2ints := kit.List2Ints(output)
			if !reflect.DeepEqual(out2ints, tt.expected) {
				t.Errorf("reverseList(%v)=%v, expected=%v", tt.input, out2ints, tt.expected)
			}
		})
	}
}
