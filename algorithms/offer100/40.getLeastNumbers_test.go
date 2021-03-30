package offer100

import (
	"reflect"
	"testing"
)

// run: go test -run -v Test_getLeastNumbers
func Test_getLeastNumbers(t *testing.T) {
	cases := []CaseEntry{
		{
			Name: "x1",
			Input: []interface{}{
				[]int{3, 2, 1}, 2,
			},
			Expected: []int{2, 1},
		},
		{
			Name: "x2",
			Input: []interface{}{
				[]int{0, 1, 2, 1}, 1,
			},
			Expected: []int{0},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			nums := c.Input.([]interface{})[0].([]int)
			k := c.Input.([]interface{})[1].(int)

			output := getLeastNumbers(nums, k)
			if !reflect.DeepEqual(output, c.Expected.([]int)) {
				t.Errorf("getLeastNumbers(%v, %d)=%v, expected=%v", nums, k, output, c.Expected)
			}
		})
	}
}
