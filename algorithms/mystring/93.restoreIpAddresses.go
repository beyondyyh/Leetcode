package mystring

import (
	"strconv"
	"strings"
)

// 暴力法，非常暴力，牛逼克拉斯
func restoreIpAddresses(s string) []string {
	var res []string
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				for d := 1; d < 4; d++ {
					if a+b+c+d == len(s) {
						n1, _ := strconv.Atoi(s[:a])
						n2, _ := strconv.Atoi(s[a : a+b])
						n3, _ := strconv.Atoi(s[a+b : a+b+c])
						n4, _ := strconv.Atoi(s[a+b+c:])
						if n1 <= 255 && n2 <= 255 && n3 <= 255 && n4 <= 255 {
							var arr []string
							arr = append(arr, strconv.Itoa(n1), strconv.Itoa(n2), strconv.Itoa(n3), strconv.Itoa(n4))
							IP := strings.Join(arr, ".")
							if len(IP) == len(s)+3 {
								res = append(res, IP)
							}
						}
					}
				}
			}
		}
	}
	return res
}
