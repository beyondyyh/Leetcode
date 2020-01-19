package myarray

import "testing"

// run: go test -v 287.*
func Test_findDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "dup is 2",
			input:    []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "dup is 3",
			input:    []int{2, 1, 3, 3, 4},
			expected: 3,
		},
		{
			name:     "dup is 6",
			input:    []int{1, 4, 6, 6, 6, 2, 3},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.input); got != tt.expected {
				t.Errorf("findDuplicate(%v)=%v, expected=%v", tt.input, got, tt.expected)
			}
		})
	}
}
