package myarray

func moveZeros1(nums []int) {
	j := 0 // 非零元素的下标
	for i := 0; i < len(nums); i++ {
		// 第一次遍历的时候，j指针记录非0的个数，只要是非0的统统都赋给nums[j]
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	// 非0元素统计完了，剩下的都是0了
	// 所以第二次遍历把末尾的元素都赋为0即可
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

// 一次遍历，借助快排思想
func moveZeroes2(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		// 当前元素!=0，就把其交换到左边，等于0的交换到右边
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
}
