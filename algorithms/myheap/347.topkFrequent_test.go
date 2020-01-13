package myheap

import (
	"reflect"
	"testing"
)

type (
	entry347 struct {
		name     string
		input    entry347input
		expected []int
	}
	entry347input struct {
		nums []int
		k    int
	}
)

// run: go test -v base.go 347.*
func Test_topkFrequent(t *testing.T) {
	cases := []entry347{
		{
			name: "x1",
			input: entry347input{
				nums: []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 4, 4},
				k:    2,
			},
			expected: []int{1, 4},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := topKFrequent(tt.input.nums, tt.input.k); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("topKFrequent(%v, %d)=%v, expected=%v", tt.input.nums, tt.input.k, output, tt.expected)
			}
		})
	}
}
