package mystring

import "testing"

// run: go test -v 7.*
func Test_reverse(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "123->321",
			input:    123,
			expected: 321,
		},
		{
			name:     "-120->-21",
			input:    -120,
			expected: -21,
		},
		{
			name:     "MaxInt溢出",
			input:    1234567899999,
			expected: 0,
		},
		{
			name:     "MinInt溢出",
			input:    -1234567899999,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if output := reverse(tt.input); output != tt.expected {
				t.Errorf("reverse(%v)=%v expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}
