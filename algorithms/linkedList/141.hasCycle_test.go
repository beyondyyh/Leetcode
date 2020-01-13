package linkedList

import (
	"testing"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

type (
	entry141 struct {
		name     string
		input    entry141input
		expected bool
	}
	entry141input struct {
		ints []int
		pos  int
	}
)

// run: go test -v base.go 141.*
func Test_hasCycle(t *testing.T) {
	cases := []entry141{
		{
			name: "x1",
			input: entry141input{
				ints: []int{},
				pos:  -1,
			},
			expected: false,
		},
		{
			name: "x2",
			input: entry141input{
				ints: []int{1},
				pos:  -1,
			},
			expected: false,
		},
		{
			name: "x3",
			input: entry141input{
				ints: []int{1, 2, 3, 4},
				pos:  1,
			},
			expected: true,
		},
		{
			name: "x4",
			input: entry141input{
				ints: []int{1, 2, 3, 4},
				pos:  3,
			},
			expected: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			head := kit.Ints2ListWithCycle(tt.input.ints, tt.input.pos)
			if output := hasCycle(head); output != tt.expected {
				t.Errorf("hasCycle(%v)=%t, expected=%t", kit.List2Ints(head, 10), output, tt.expected)
			}
		})
	}
}
