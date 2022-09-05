package mystring

// 1. 暴力解法: 列出子串, 筛选出长度最大值
// 时间复杂度：O(n^3)
// 空间复杂度：O(1)
func longestPalindrome1(s string) string { // {{{
	if s == "" {
		return ""
	}
	if isPalindrome(s) {
		return s
	}
	var max int
	var res string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substr := s[i:j]
			// fmt.Println(substr)
			if isPalindrome(substr) {
				if len(substr) > max {
					max = len(substr)
					res = substr
				}
			}
		}
	}
	if res == "" {
		return s[0:1]
	}
	return res
} // }}}

func isPalindrome(s string) bool { // {{{
	if len(s) < 1 {
		return false
	}
	if len(s) == 2 && s[0] == s[1] {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
} // }}}

// 2. 中心扩展法
// 我们观察到回文中心的两侧互为镜像。因此，回文可以从它的中心展开，并且只有 2n - 1 个这样的中心。
// 你可能会问，为什么会是 2n - 1 个，而不是 n 个中心？
// 因为回文的中心要区分单双。
// 假如回文的中心为 双数，例如 abba，那么可以划分为 ab bb ba，对于n长度的字符串，这样的划分有 n-1 种。
// 假为回文的中心为 单数，例如 abcd, 那么可以划分为 a b c d， 对于n长度的字符串，这样的划分有 n 种。
// 对于 n 长度的字符串，我们其实不知道它的回文串中心倒底是单数还是双数，所以我们要对这两种情况都做遍历，也就是 n+(n-1) = 2n - 1，所以时间复杂度为 O(n)。
// 当中心确定后，我们要围绕这个中心来扩展回文，那么最长的回文可能是整个字符串，所以时间复杂度为 O(n)。
// 所以总时间复杂度为 O(n^2)
func longestPalindrome2(s string) string {
	if len(s) < 1 {
		return ""
	}

	var start, end int
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // 一个元素为中心进行扩展
		len2 := expandAroundCenter(s, i, i+1) // 两个元素为中心
		length := max(len1, len2)
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
		// fmt.Printf("i:%d len1:%d len2:%d length:%d sub:%s\n", i, len1, len2, length, s[start:end+1])
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int { // {{{
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
} // }}}

func max(a, b int) int { // {{{
	if a >= b {
		return a
	} else {
		return b
	}
} // }}}
