package mystring

func titleToNumber(s string) int {
	ans := 0
	for _, c := range s {
		num := c - 'A' + 1
		ans = ans*26 + int(num)
	}
	return ans
}
