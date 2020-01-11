package tree

import (
	"reflect"
	"sort"
	"testing"
)

// run: go test -v 22.*
func Test_generateParenthesis(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected []string
	}{
		{
			name:     "x1",
			input:    1,
			expected: []string{"()"},
		},
		{
			name:     "x2",
			input:    2,
			expected: []string{"()()", "(())"},
		},
		{
			name:     "x3",
			input:    3,
			expected: []string{"()(())", "()()()", "((()))", "(()())", "(())()"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output := generateParenthesis(tt.input)
			sort.Strings(output)
			sort.Strings(tt.expected)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("generateParenthesis(%d)=\"%v\", expected=\"%v\"", tt.input, output, tt.expected)
			}
		})
	}
}
