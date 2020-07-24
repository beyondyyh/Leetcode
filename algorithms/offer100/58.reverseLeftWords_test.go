package offer100

import (
	"testing"
)

// run: go test -run Test_reverseLeftWords
func Test_reverseLeftWords(t *testing.T) {
	cases := []CaseEntry{
		{
			Name: "x1",
			Input: []interface{}{
				"abcdefg", 2,
			},
			Expected: "cdefgab",
		},
		{
			Name: "x2",
			Input: []interface{}{
				"lrloseumgh", 6,
			},
			Expected: "umghlrlose",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			s := c.Input.([]interface{})[0].(string)
			k := c.Input.([]interface{})[1].(int)
			if output := reverseLeftWords1(s, k); output != c.Expected.(string) {
				t.Errorf("reverseLeftWords1(%s, %d)=%s, expected=%s", s, k, output, c.Expected.(string))
			}

			if output := reverseLeftWords2(s, k); output != c.Expected.(string) {
				t.Errorf("reverseLeftWords2(%s, %d)=%s, expected=%s", s, k, output, c.Expected.(string))
			}
		})
	}
}
