package mystring

import (
	"Leetcode/algorithms/kit"
	"testing"
)

// run: go test -v 13.*
func Test_romanToInt(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    "III",
			Expected: 3,
		},
		{
			Name:     "x2",
			Input:    "IV",
			Expected: 4,
		},
		{
			Name:     "x3",
			Input:    "IX",
			Expected: 9,
		},
		{
			Name:     "x4",
			Input:    "LVIII",
			Expected: 58,
		},
		{
			Name:     "x5",
			Input:    "MCMXCIV",
			Expected: 1994,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if output := romanToInt1(c.Input.(string)); output != c.Expected.(int) {
				t.Errorf("romanToInt1(%s)=%d, expected=%d", c.Input, output, c.Expected)
			}

			if output := romanToInt2(c.Input.(string)); output != c.Expected.(int) {
				t.Errorf("romanToInt2(%s)=%d, expected=%d", c.Input, output, c.Expected)
			}
		})
	}
}
