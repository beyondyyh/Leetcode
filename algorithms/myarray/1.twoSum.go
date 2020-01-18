package myarray

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for k, v := range nums {
		diff := target - v
		if kk, exists := hash[diff]; exists {
			return []int{kk, k}
		}
		hash[v] = k
	}
	return []int{}
}
