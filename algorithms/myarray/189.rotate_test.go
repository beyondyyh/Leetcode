package myarray

import (
	"reflect"
	"testing"
)

type (
	entry189input struct {
		nums []int
		k    int
	}
	entry189 struct {
		name     string
		input    entry189input
		expected []int
	}
)

// run: go test -v 121.*
func Test_rotate(t *testing.T) {
	cases := []entry189{
		{
			name: "x1",
			input: entry189input{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "x2",
			input: entry189input{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expected: []int{3, 99, -1, -100},
		},
		{
			name: "x3",
			input: entry189input{
				nums: []int{1, -1},
				k:    1,
			},
			expected: []int{-1, 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t1, t2 := c, c
			rotate1(t1.input.nums, t1.input.k)
			if !reflect.DeepEqual(t1.input.nums, t1.expected) {
				t.Errorf("rotate1(%v,%d)=%v, expected=%v", t1.input.nums, t1.input.k, t1.input.nums, t1.expected)
			}

			rotate2(t2.input.nums, t2.input.k)
			if !reflect.DeepEqual(t2.input.nums, t2.expected) {
				t.Errorf("rotate2(%v,%d)=%v, expected=%v", t2.input.nums, t2.input.k, t2.input.nums, t2.expected)
			}
		})
	}
}
