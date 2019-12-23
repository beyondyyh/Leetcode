package main

func findDuplicate(nums []int) int {
	// slow, fast := 0, 0
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	// 让before，after分别指向链表开始节点，相遇节点
	before, after := 0, slow
	for before != after {
		before, after = nums[before], nums[after]
	}
	return before
}
