package main

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "123->321",
			input: 123,
			want:  321,
		},
		{
			name:  "-120->-21",
			input: -120,
			want:  -21,
		},
		{
			name:  "MaxInt溢出",
			input: 1234567899999,
			want:  0,
		},
		{
			name:  "MinInt溢出",
			input: -1234567899999,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.input); got != tt.want {
				t.Errorf("reverse(%v)=%v want=%v", tt.input, got, tt.want)
			}
		})
	}
}
