package main

import (
	"fmt"
	"math"
)

/*
给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
找到所有出现两次的元素。
你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
*/

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(arr))
}

// 关键点：1 <= a[i] <= n，n为数组长度，所以a[i]可以转换为数组下标~~~~~
// 在输入的数组中用数字的正负数来标识该位置所对应的数字是否已经出现过；
// 遍历数组，对对应位置的数字取相反数，如果已经是负数则说明前面出现过了。
func findDuplicates(arr []int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		num := int(math.Abs(float64(arr[i])))
		if arr[num-1] > 0 {
			arr[num-1] *= -1
		} else {
			res = append(res, num)
		}
		fmt.Printf("num:%v arr:%v res:%v\n", num, arr, res)
	}
	return res
}

// output
// num:4 arr:[4 3 2 -7 8 2 3 1] res:[]
// num:3 arr:[4 3 -2 -7 8 2 3 1] res:[]
// num:2 arr:[4 -3 -2 -7 8 2 3 1] res:[]
// num:7 arr:[4 -3 -2 -7 8 2 -3 1] res:[]
// num:8 arr:[4 -3 -2 -7 8 2 -3 -1] res:[]
// num:2 arr:[4 -3 -2 -7 8 2 -3 -1] res:[2]
// num:3 arr:[4 -3 -2 -7 8 2 -3 -1] res:[2 3]
// num:1 arr:[-4 -3 -2 -7 8 2 -3 -1] res:[2 3]
