package linkedList

import (
	"testing"
)

func Test_deleteNode(t *testing.T) {
	list1 := buildList(200) // 返回的是头结点
	deleteNode(list1)
	dump(list1) // 202->203->204->205
}
