package main

import (
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "x1",
			input:    "/../",
			expected: "/",
		},
		{
			name:     "x2",
			input:    "/home//foo/",
			expected: "/home/foo",
		},
		{
			name:     "x3",
			input:    "/a/./b/../../c/",
			expected: "/c",
		},
		{
			name:     "x4",
			input:    "/a/../../b/../c//.//",
			expected: "/c",
		},
		{
			name:     "x5",
			input:    "/a//b////c/d//././/..",
			expected: "/a/b/c",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := simplifyPath(tt.input); output != tt.expected {
				t.Errorf("simplifyPath(%s)=%s, expected=%s", tt.input, output, tt.expected)
			}
		})
	}
}
