package kit

// ListNode 单链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	dummy := &ListNode{Val: -1}
	curr := dummy
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return dummy.Next
}

// List2Ints convert list to []int
// limit 链表深度限制，防止环状链表无限循环，设置limit超过会panic，默认20
func List2Ints(head *ListNode, limits ...int) []int {
	if head == nil {
		return []int{}
	}

	limit := 20
	if len(limits) > 0 && limits[0] > 0 {
		limit = limits[0]
	}

	var times int
	var res []int
	for head != nil {
		times++
		if times > limit {
			// panic(fmt.Sprintf("depth exceeds maximum limit %d", limit))
			break
		}
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos <= 0 or pos >= len(nums), no cycle
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos <= 0 || pos >= len(nums) {
		return head
	}

	// find the pos-indexed node
	curr := head
	for pos > 0 {
		curr = curr.Next
		pos--
	}
	// the remaining nodes
	tail := curr
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = curr

	return head
}

// GetNodeWith returns the first node with val
func (head *ListNode) GetNodeWith(val int) *ListNode {
	curr := head
	for curr != nil {
		if curr.Val == val {
			break
		}
		curr = curr.Next
	}
	return curr
}
