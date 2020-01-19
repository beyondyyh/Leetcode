package myarray

import (
	"testing"
)

// run: go test -v 26.*
func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "case 1",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "case 2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.input); got != tt.expected {
				t.Errorf("removeDuplicates(%v)=%v, expected=%v", tt.input, got, tt.expected)
			}
		})
	}
}
