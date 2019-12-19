package tree

import (
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	n := 3
	res := generateParenthesis(n)
	t.Logf("when n=%d res=%v", n, res)
}
