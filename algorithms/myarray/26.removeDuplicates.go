package myarray

// 双指针法
// i慢指针，j快指针，i从0开始j从1开始，j先走
// 如果nums[j] = nums[i]说明是重复项，i不动，j继续往前走
// 直到nums[j] != nums[i]说明重复运动结束，i走一步，同时把nums[j]赋值给num[i]
// 时间复杂度：O(n)
func removeDuplicates1(nums []int) int {
	L := len(nums)
	if L == 0 {
		return 0
	}

	i := 0 // 慢指针i
	for j := 1; j < L; j++ {
		// 快指针j先走，直到重复运动结束，把nums[j]赋值给i+1，同时慢指针i走一步
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}

func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	i := 0
	for j := 1; j < n; j++ {
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}
