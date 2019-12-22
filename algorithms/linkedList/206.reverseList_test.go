package linkedList

import (
	"fmt"
	"testing"
)

// run: go test -v ListNode.go 206.*
func Test_reverseList(t *testing.T) {
	list1 := buildList(0)
	// dump(list1)
	fmt.Println("**********")

	list2 := reverseList(list1)
	dump(list2)
}
