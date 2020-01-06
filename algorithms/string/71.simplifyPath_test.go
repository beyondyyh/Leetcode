package main

import (
	"testing"
)

type entry71 struct {
	name     string
	input    string
	expected string
}

// build_simplifyPath_case return test cases
func build_simplifyPath_case() []entry71 {
	cases := []entry71{
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
	return cases
}

// run: go test -v 71.*
func Test_simplifyPath(t *testing.T) {
	for _, tt := range build_simplifyPath_case() {
		t.Run(tt.name, func(t *testing.T) {
			if output := simplifyPath(tt.input); output != tt.expected {
				t.Errorf("simplifyPath(%s)=%s, expected=%s", tt.input, output, tt.expected)
			}
		})
	}
}

func Test_cleanFilepath(t *testing.T) {
	for _, tt := range build_simplifyPath_case() {
		t.Run(tt.name, func(t *testing.T) {
			if output := cleanFilepath(tt.input); output != tt.expected {
				t.Errorf("cleanFilepath(%s)=%s, expected=%s", tt.input, output, tt.expected)
			}
		})
	}
}
