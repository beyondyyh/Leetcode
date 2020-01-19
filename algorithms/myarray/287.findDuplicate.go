package myarray

import "fmt"

func findDuplicate(nums []int) int {
	fmt.Println(nums)
	slow, fast := nums[0], nums[nums[0]]
	fmt.Printf("slow:%d fast:%d\n", slow, fast)
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
		fmt.Printf("slow:%d fast:%d\n", slow, fast)
	}

	fmt.Println("------------------")
	// 让slow，fast分别指向链表开始节点，相遇节点
	slow = 0 // fast此时就是指向相遇节点
	fmt.Printf("slow:%d fast:%d\n", slow, fast)
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
		fmt.Printf("slow:%d fast:%d\n", slow, fast)
	}
	fmt.Println("**************")
	return slow
}

// [1 3 4 2 2]
// slow:1 fast:3
// slow:3 fast:4
// slow:2 fast:4
// slow:4 fast:4
// ------------------
// slow:0 fast:4
// slow:1 fast:2
// slow:3 fast:4
// slow:2 fast:2
