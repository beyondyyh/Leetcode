package myqueue

// Queue for int array
type Queue struct {
	nums []int
}

// NewQueue returns *myqueue.Queue
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
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

// IsEmpty returns the q is empty or not
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
