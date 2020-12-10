package linkedList

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

type (
	entry92 struct {
		name     string
		input    entry92input
		expected []int
	}
	entry92input struct {
		ints []int
		m    int
		n    int
	}
)

// run: go test -v base.go 92.*
func Test_reverseBetween(t *testing.T) {
	cases := []entry92{
		{
			name:     "x1",
			input:    entry92input{ints: []int{1, 2, 3, 4, 5}, m: 2, n: 4},
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			name:     "x2",
			input:    entry92input{ints: []int{1, 2, 3, 4, 5}, m: 1, n: 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range cases {
		head := kit.Ints2List(tc.input.ints)
		output := kit.List2Ints(reverseBetween(head, tc.input.m, tc.input.n))
		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("reverseBetween(%v, %d, %d)=%v, expected=%v", tc.input.ints, tc.input.m, tc.input.n, output, tc.expected)
		}
	}
}
