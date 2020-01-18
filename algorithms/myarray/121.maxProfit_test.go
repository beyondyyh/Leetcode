package myarray

import "testing"

// run: go test -v 121.*
func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "x1",
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "x2",
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "x3",
			input:    []int{10, 2, 9, 1, 2, 1, 3, 1},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if output := maxProfit(tt.input); output != tt.expected {
				t.Errorf("maxProfit(%v)=%d, expected=%d\n", tt.input, output, tt.expected)
			}
		})
	}
}
