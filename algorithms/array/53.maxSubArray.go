package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 3, -1, 1, -3}
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	var ans int = nums[0]
	var sum int = 0
	for i, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		fmt.Printf("i:%d num:%d sum:%d ans:%d\n", i, num, sum, ans)

		if sum > ans {
			ans = sum
		}
	}
	return ans
}

// i:0 num:-2 sum:-2 ans:-2
// i:1 num:3 sum:3 ans:-2
// i:2 num:-1 sum:2 ans:3
// i:3 num:1 sum:3 ans:3
// i:4 num:-3 sum:0 ans:3
// 3
