package mybit

// 交换律：a ^ b ^ c <=> a ^ c ^ b
// 任何数于0异或为任何数 0 ^ n => n
// 相同的数异或为0: n ^ n => 0
// var a = [2,3,2,4,4]
// 2 ^ 3 ^ 2 ^ 4 ^ 4等价于 2 ^ 2 ^ 4 ^ 4 ^ 3 => 0 ^ 0 ^3 => 3
func singleNumber(nums []int) int {
	var ans int
	for _, num := range nums {
		ans ^= num
	}

	return ans
}
