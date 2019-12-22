package linkedList

import "fmt"

// ListNode 单链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func insert(val int, head *ListNode) *ListNode {
	var node *ListNode = &ListNode{Val: val}
	node.Next = head
	return node
}

func dump(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("cur.Val:%d cur.Next:%v\n", cur.Val, cur.Next)
		cur = cur.Next
	}
}

func display(head *ListNode) []interface{} {
	cur := head
	var items []interface{}
	for cur != nil {
		items = append(items, cur.Val)
		cur = cur.Next
	}
	items = append(items, "nil")
	return items
}

/**
 * 构造List
 *	when step = 0: 1->2->3->4->5
 *	when step = 100: 101->102->103->104->105
 */
func buildList(step int) *ListNode {
	head := new(ListNode)
	head.Val = step + 1
	ln2 := new(ListNode)
	ln2.Val = step + 2
	ln3 := new(ListNode)
	ln3.Val = step + 3
	/*
		ln4 := new(ListNode)
		ln4.Val = step + 4
		ln5 := new(ListNode)
		ln5.Val = step + 5
	*/
	head.Next = ln2
	ln2.Next = ln3
	// ln3.Next = ln4
	// ln4.Next = ln5
	return head
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	dummy := &ListNode{Val: -1}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}

// List2Ints convert list to []int
func List2Ints(head *ListNode) []int {
	var items []int
	cur := head
	for cur != nil {
		items = append(items, cur.Val)
		cur = cur.Next
	}
	// items = append(items, "nil")
	return items
}
