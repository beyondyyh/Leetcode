package myheap

import "testing"

type (
	entry215 struct {
		name     string
		input    entry215input
		expected int
	}
	entry215input struct {
		nums []int
		k    int
	}
)

// run: go test -v base.go 215.*
func Test_findKthLargest(t *testing.T) {
	cases := []entry215{
		{
			name: "x1",
			input: entry215input{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			expected: 5,
		},
		{
			name: "x2",
			input: entry215input{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    4,
			},
			expected: 3,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := findKthLargest(tt.input.nums, tt.input.k); output != tt.expected {
				t.Errorf("findKthLargest(%v,%d)=%d, expected=%d", tt.input.nums, tt.input.k, output, tt.expected)
			}
		})
	}
}
