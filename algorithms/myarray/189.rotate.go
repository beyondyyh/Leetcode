package myarray

// 方法一：引入额外数组newNums，遍历将nums[i]赋值给newNums[(i+k)%n]
// 时间复杂度：O(n)，n为数组长度
// 空间复杂度：O(n)
func rotate1(nums []int, k int) {
	n := len(nums)
	newNums := make([]int, n)
	for i := 0; i < n; i++ {
		newNums[(i+k)%n] = nums[i]
	}
	copy(nums, newNums) // func copy(dst, src []Type) int
}

// 方法二：多次翻转数组，先将nums翻转，然后再分别翻转前k个 和 后n-k个
// 时间复杂度：n/2 + k/2 + (n-k)/2，故为O(n)
// 空间复杂度：没有引入额外存储，故为O(1)
func rotate2(nums []int, k int) {
	k %= len(nums)    // 这个妙，巧妙地避免了数组越界的情况
	reverse(nums)     // [7,6,5,4,3,2,1]
	reverse(nums[:k]) // [5,6,7,4,3,2,1]
	reverse(nums[k:]) // [5,6,7,1,2,3,4]
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
