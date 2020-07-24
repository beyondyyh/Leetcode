package offer100

// 利用字符串切片
func reverseLeftWords1(s string, n int) string {
	return s[n:] + s[:n]
}

// 遍历，巧妙利用mod
func reverseLeftWords2(s string, n int) string {
	var ans string
	for i := n; i < len(s)+n; i++ {
		ans += string(s[i%len(s)])
	}
	return ans
}
