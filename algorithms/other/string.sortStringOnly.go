package main

import (
	"fmt"
	"sort"
)

// 题目：给定一个字符串，带数字字母和其他特殊字符，要求只给字母排序，其他字符串位置保持不变
// 例如：input: cb12$%a2A output: Aa12$%b2c
func main() {
	str := "09cbAa12d98xYz~#"
	// str = "cb12$%a2A"
	fmt.Printf("origin str: %s\n", str)
	fmt.Println(sortStringOnly(str))
}

func sortStringOnly(str string) string {
	var (
		letter, other []string
		flags         []bool
	)
	for _, char := range []byte(str) {
		// A-Z || a-z
		if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
			letter = append(letter, string(char))
			flags = append(flags, true)
		} else {
			other = append(other, string(char))
			flags = append(flags, false)
		}
	}

	sort.Strings(letter)

	var i, j int
	var res string
	for _, flag := range flags {
		if flag {
			res += letter[i]
			i++
		} else {
			res += other[j]
			j++
		}
	}
	return res
}
