package main

import "testing"

// run: go test -v 287.*
func Test_findDuplicate(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "dup is 2",
			input: []int{1, 3, 4, 2, 2},
			want:  2,
		},
		{
			name:  "dup is 3",
			input: []int{2, 1, 3, 3, 4},
			want:  3,
		},
		{
			name:  "dup is 6",
			input: []int{1, 4, 6, 6, 6, 2, 3},
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.input); got != tt.want {
				t.Errorf("findDuplicate(%v)=%v, want=%v", tt.input, got, tt.want)
			}
		})
	}
}
