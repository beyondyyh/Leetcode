package linkedList

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

type (
	entry19 struct {
		name     string
		input    entry19input
		expected []int
	}
	entry19input struct {
		ints []int
		n    int
	}
)

// go test -v base.go 19.*
func build_removeNthFromEndCase() []entry19 {
	cases := []entry19{
		{
			name: "x1",
			input: entry19input{
				ints: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			expected: []int{1, 2, 3, 5},
		},
		{
			name: "x2",
			input: entry19input{
				ints: []int{1, 2, 3, 4, 5},
				n:    5,
			},
			expected: []int{2, 3, 4, 5},
		},
		{
			name: "x3",
			input: entry19input{
				ints: []int{1, 2, 3, 4, 5},
				n:    6,
			},
			expected: []int{2, 3, 4, 5},
		},
	}
	return cases
}

// run: go test -v base.go 19.*
func Test_removeNthFromEnd1(t *testing.T) {
	for _, tt := range build_removeNthFromEndCase() {
		t.Run(tt.name, func(t *testing.T) {
			head := kit.Ints2List(tt.input.ints)
			output := removeNthFromEnd1(head, tt.input.n)
			out2ints := kit.List2Ints(output)
			if !reflect.DeepEqual(out2ints, tt.expected) {
				t.Errorf("removeNthFromEnd1(%v, %d)=%v, expected=%v", tt.input.ints, tt.input.n, out2ints, tt.expected)
			}
		})
	}
}

func Test_removeNthFromEnd2(t *testing.T) {
	for _, tt := range build_removeNthFromEndCase() {
		t.Run(tt.name, func(t *testing.T) {
			head := kit.Ints2List(tt.input.ints)
			output := removeNthFromEnd2(head, tt.input.n)
			out2ints := kit.List2Ints(output)
			if !reflect.DeepEqual(out2ints, tt.expected) {
				t.Errorf("removeNthFromEnd2(%v, %d)=%v, expected=%v", tt.input.ints, tt.input.n, out2ints, tt.expected)
			}
		})
	}
}
