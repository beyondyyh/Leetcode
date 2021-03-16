package other

import "fmt"

// 快速排序是不稳定排序，时间复杂度O(Nlog(N))
// quickSort 快排，会延伸很多变种
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := nums[0]
	head, tail := 0, len(nums)-1
	for i := 1; i <= tail; {
		fmt.Printf("i:%d, head:%d, tail:%d, mid:%d, nums[i]:%d, nums:%v\n",
			i, head, tail, mid, nums[i], nums)
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	// nums[head] = mid
	quickSort(nums[:head])
	quickSort(nums[head+1:])
	return nums
}
