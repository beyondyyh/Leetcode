package linkedList

// DListNode 双向链表基本定义
// Key, Val 可以是任意结构
type DListNode struct {
	Key  int
	Val  int
	Next *DListNode
	Prev *DListNode
}

// 实现如下funcs
// 链表长度 Len
// 头部插入 Insert
// 尾部插入 Append
// 移除元素 Remove
