package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

// 无重复字符的最长子串, https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 1. 暴力解法是遍历取出所有子串，然后写一个allUnique方法check字符串是否不重复，利用map，在不重复的字串中找出最长的
// 2. 滑动窗口
// 3. 利用两数之和的方法
func lengthOfLongestSubstring(s string) int {
	dict := make(map[byte]int)
	res, l := 0, 0
	for i, char := range []byte(s) {
		if it, ok := dict[char]; ok && l <= it {
			l = it + 1
		}
		fmt.Printf("dict:%v char:%s i:%d l:%d\n", dict, string(char), i, l)

		// 求最大值
		if i-l+1 > res {
			res = i - l + 1
		}

		dict[char] = i
	}
	return res
}
