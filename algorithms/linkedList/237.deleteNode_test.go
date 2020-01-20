package linkedList

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

type (
	entry237 struct {
		name     string
		input    entry237input
		expected []int
	}
	entry237input struct {
		ints []int
		val  int
	}
)

// run: go test -v base.go 237.*
func Test_deleteNode(t *testing.T) {
	cases := []entry237{
		{
			name: "x1",
			input: entry237input{
				ints: []int{4, 5, 1, 9},
				val:  5,
			},
			expected: []int{4, 1, 9},
		},
		{
			name: "x2",
			input: entry237input{
				ints: []int{4, 5, 1, 9},
				val:  1,
			},
			expected: []int{4, 5, 9},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			head := kit.Ints2List(tt.input.ints)
			node := head.GetNodeWith(tt.input.val)
			deleteNode(node)
			output := kit.List2Ints(head)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("origin Link=%v after deleteNode(%d)=%v, expected=%v", tt.input.ints, tt.input.val, output, tt.expected)
			}
		})
	}
}
