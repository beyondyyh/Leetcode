package offer100

import (
	"reflect"
	"testing"
)

type (
	entry239input struct {
		nums []int
		k    int
	}
	entry239 struct {
		name     string
		input    entry239input
		expected []int
	}
)

// run: go test -run Test_maxSlidingWindow
func Test_maxSlidingWindow(t *testing.T) {
	cases := []entry239{
		{
			name: "x1",
			input: entry239input{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "x2",
			input: entry239input{
				nums: []int{1},
				k:    1,
			},
			expected: []int{1},
		},
		{
			name: "x3",
			input: entry239input{
				nums: []int{1, -1},
				k:    1,
			},
			expected: []int{1, -1},
		},
		{
			name: "x4",
			input: entry239input{
				nums: []int{9, 11},
				k:    2,
			},
			expected: []int{11},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := maxSlidingWindow(c.input.nums, c.input.k)
			if !reflect.DeepEqual(output, c.expected) {
				t.Errorf("maxSlidingWindow(%v,%d)=%v, expected=%v", c.input.nums, c.input.k, output, c.expected)
			}
		})
	}
}
