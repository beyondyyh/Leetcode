package mystring

import "strings"

// 最长公共前缀, https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	var res = ""
	if len(strs) == 0 {
		return res
	}

	var prefix string = strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 { // 遍历，如果不包含，就-1即向前移一位
			prefix = prefix[0 : len(prefix)-1]
		}
	}
	return prefix
}
