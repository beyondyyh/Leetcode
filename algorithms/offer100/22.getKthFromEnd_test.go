package offer100

import (
	"Leetcode/algorithms/kit"
	"reflect"
	"testing"
)

// run: go test -run Test_getKthFromEnd
func Test_getKthFromEnd(t *testing.T) {
	cases := []CaseEntry{
		{
			Name: "x1",
			Input: []interface{}{
				[]int{1, 2, 3, 4, 5}, 2,
			},
			Expected: []int{4, 5},
		},
		{
			Name: "x1",
			Input: []interface{}{
				[]int{1, 2, 3, 4, 5}, 4,
			},
			Expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			head := kit.Ints2List(c.Input.([]interface{})[0].([]int))
			k := c.Input.([]interface{})[1].(int)

			out2Ints := kit.List2Ints(getKthFromEnd(head, k))
			exp2Ints := c.Expected.([]int)
			if !reflect.DeepEqual(out2Ints, exp2Ints) {
				t.Errorf("getKthFromEnd(%v, %d)=%v, expected=%v", c.Input.([]interface{})[0].([]int), k, out2Ints, exp2Ints)
			}
		})
	}
}
