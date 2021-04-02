package offer100

import (
	"testing"
)

var matrix = [][]int{
	{1, 4, 7, 11, 15},
	{2, 5, 8, 12, 19},
	{3, 6, 9, 16, 22},
	{10, 13, 14, 17, 24},
	{18, 21, 23, 26, 30},
}

type entry04input struct {
	name     string
	input    int
	expected bool
}

// run: go test -run Test_findNumberIn2DArray
func Test_findNumberIn2DArray(t *testing.T) {
	cases := []entry04input{
		{
			name:     "x1",
			input:    5,
			expected: true,
		},
		{
			name:     "x2",
			input:    18,
			expected: true,
		},
		{
			name:     "x3",
			input:    20,
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := findNumberIn2DArray(matrix, c.input)
			if output != c.expected {
				t.Errorf("findNumberIn2DArray(%d)=%t, expected=%t", c.input, output, c.expected)
			}
		})
	}
}
