package linkedList

import "fmt"

type Element interface{}

type Node struct {
	Data Element
	Next *Node
}

type List struct {
	HeadNode *Node // 头节点,nil
}

// dump链表
func (l *List) Dump() {
	fmt.Println("*************")
	cur := l.HeadNode
	for cur != nil {
		fmt.Printf("Data:%v Next:%v Ptr:%v\n", cur.Data, cur.Next, *cur)
		cur = cur.Next
	}
	fmt.Println("*************")
}

// 链表长度
func (l *List) Length() int {
	cur := l.HeadNode
	cnt := 0
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	return cnt
}

// 判断是否为空链表
func (l *List) IsEmpty() bool {
	if l.HeadNode == nil {
		return true
	}
	return false
}

// 头部插入
func (l *List) Insert(e Element) {
	node := &Node{Data: e}
	if l.IsEmpty() {
		l.HeadNode = node
	} else {
		node.Next = l.HeadNode
		l.HeadNode = node
	}
}

// 尾部追加
func (l *List) Append(e Element) {
	node := &Node{Data: e}
	if l.IsEmpty() {
		l.HeadNode = node
	} else {
		cur := l.HeadNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

// 删除不重复单链表中的某个元素
func (l *List) Remove(e Element) {
	cur := l.HeadNode
	if cur.Data == e {
		l.HeadNode = cur.Next
	} else {
		for cur.Next != nil {
			if cur.Next.Data == e {
				cur.Next = cur.Next.Next
			} else {
				cur = cur.Next
			}
		}
	}
}

// 是否包含某个元素
func (l *List) Contain(e Element) bool {
	cur := l.HeadNode
	for cur != nil {
		if cur.Data == e {
			return true
		}
		cur = cur.Next
	}
	return false
}

// 在指定位置删除元素
func (l *List) RemoveAtIndex(i int) {
	if i >= l.Length() {
		panic("i lt l.Length()")
	}

	cur := l.HeadNode
	if i <= 0 {
		l.HeadNode = cur.Next
	} else {
		count := 0
		for count != (i-1) && cur.Next != nil {
			count++
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}
}

// 在指定位置添加元素
func (l *List) InsertAtIndex(i int, e Element) {
	if i <= 0 {
		l.Insert(e)
	} else if i > l.Length() {
		l.Append(e)
	} else {
		cur := l.HeadNode
		count := 0
		for count < (i - 1) {
			cur = cur.Next
			count++
		}
		// 当循环退出时，cur指向index-1的位置
		node := &Node{Data: e}
		node.Next = cur.Next
		cur.Next = node
	}
}

////////
// 链表反转
func (l *List) Reverse() *Node {
	cur := l.HeadNode
	var pre *Node = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
