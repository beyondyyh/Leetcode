package linkedList

import (
	"fmt"
	"testing"
)

// run: go test -v ListNode.go 2.*
func Test_addTwoNumbers(t *testing.T) {
	l1 := buildList(0)
	l2 := buildList(4)
	l3 := addTwoNumbers(l1, l2)
	fmt.Println("l1:")
	dump(l1)
	fmt.Println("l2:")
	dump(l2)
	fmt.Println("l3:")
	dump(l3)
}
