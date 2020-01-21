package tree

import (
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v 331.*
func Test_isValidSerialization(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    "9,3,4,#,#,1,#,#,2,#,6,#,#",
			Expected: true,
		},
		{
			Name:     "x2",
			Input:    "1,#",
			Expected: false,
		},
		{
			Name:     "x3",
			Input:    "9,#,#,1",
			Expected: false,
		},
		// add more test cases
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			if output := isValidSerialization1(tt.Input.(string)); output != tt.Expected.(bool) {
				t.Errorf("isValidSerialization1(%s)=%t, expected=%t", tt.Input, output, tt.Expected.(bool))
			}

			if output := isValidSerialization2(tt.Input.(string)); output != tt.Expected.(bool) {
				t.Errorf("isValidSerialization2(%s)=%t, expected=%t", tt.Input, output, tt.Expected.(bool))
			}
		})
	}
}
