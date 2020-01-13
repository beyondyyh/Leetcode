package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	fmt.Printf("original:%s ", s)
	// flag表示正负，res最终结果
	flag, res := 1, 0

	// 去掉所有首位空白，如: \t\r\n
	s = strings.TrimSpace(s)
	fmt.Printf("after trim:%s\n", s)
	if len(s) == 0 {
		return res
	}

	var i int = 0
	if s[i] == '+' {
		i++
		flag = 1
	} else if s[i] == '-' {
		i++
		flag = -1
	}

	for ; i < len(s); i++ {
		num := flag * res
		fmt.Printf("num:%v ", num)
		if num >= math.MaxInt32 {
			return math.MaxInt32
		}
		if num <= math.MinInt32 {
			return math.MinInt32
		}
		if s[i] < '0' || s[i] > '9' {
			return num
		}
		res = res*10 + int(s[i]) - '0'
	}

	num := flag * res
	if num >= math.MaxInt32 {
		return math.MaxInt32
	}
	if num <= math.MinInt32 {
		return math.MinInt32
	}
	fmt.Printf("num:%v\n%s\n", num, "*********")
	return num
}
