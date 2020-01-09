package myqueue

// Queue for int array
type Queue struct {
	nums []int
}

// NewQueue returns *Queue
func NewQueue() *Queue {
	return &Queue{nums: make([]int, 0)}
}

// Len returns the length of s
func (q *Queue) Len() int {
	return len(q.nums)
}

// Push push n to q
func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

// Pop pop the first element from q
func (q *Queue) Pop() int {
	x := q.nums[0]
	q.nums = q.nums[1:]
	return x
}

// Peek returns the front element of q, keep the original structure of q
func (q *Queue) Peek() int {
	return q.nums[0]
}

// IsEmpty returns the q is empty or not
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

// PeekAll returns the whole eles, for debug
func (q *Queue) PeekAll() []int {
	return q.nums
}
