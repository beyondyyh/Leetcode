package myarray

import (
	"reflect"
	"sort"
	"testing"
)

type (
	entry1 struct {
		name     string
		input    entry1input
		expected []int
	}
	entry1input struct {
		nums   []int
		target int
	}
)

// run: go test -v 1.*
func Test_twoSum(t *testing.T) {
	cases := []entry1{
		{
			name: "x1",
			input: entry1input{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			name: "x1",
			input: entry1input{
				nums:   []int{3, 2, 4},
				target: 8,
			},
			expected: []int{},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output := twoSum(tt.input.nums, tt.input.target)
			sort.Ints(output)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("twoSum(%v, %d)=%v, expected=%v", tt.input.nums, tt.input.target, output, tt.expected)
			}
		})
	}
}
