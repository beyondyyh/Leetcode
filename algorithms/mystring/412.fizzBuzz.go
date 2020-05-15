package mystring

import "strconv"

func fizzBuzz(n int) []string {
	zzmap := map[string]string{
		"3": "Fizz",
		"5": "Buzz",
	}

	var ans []string
	for i := 1; i <= n; i++ {
		var item string
		for ks, vs := range zzmap {
			k, _ := strconv.Atoi(ks)
			if i%k == 0 {
				item += vs
			}
		}
		// 都不能整除
		if item == "" {
			item = strconv.Itoa(i)
		}
		ans = append(ans, item)
	}

	return ans
}
