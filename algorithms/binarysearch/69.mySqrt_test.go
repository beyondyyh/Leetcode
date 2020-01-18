package binarysearch

import "testing"

// run: go test -v 69.*
func Test_mySqrt(t *testing.T) {
	cases := []struct {
		name            string
		input, expected int
	}{
		{"x0", 1, 1},
		{"x1", 2, 1},
		{"x2", 4, 2},
		{"x3", 5, 2},
		{"x4", 8, 2},
		{"x5", 9, 3},
		{"x6", 100, 10},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := mySqrt(tt.input); output != tt.expected {
				t.Errorf("mySqrt(%d)=%d, expected=%d", tt.input, output, tt.expected)
			}
		})
	}
}
