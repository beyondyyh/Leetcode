package mystring

// 思路：建立一个Map来映射符号和值，然后对字符串从后向前遍历，保留上一个字符串
// 如果当前字符代表的值大于等于上一个（右边），就加上该值；
// 如果当前字符代表的值小于上一个（右边），则减去该值。
// romanToInt1 returns int
func romanToInt1(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res, last int
	for i := len(s) - 1; i >= 0; i-- {
		curr := m[s[i]]
		flag := 1
		// 小数在大数的左边，要减去小数
		if curr < last {
			flag = -1
		}

		res += flag * curr
		last = curr
	}

	return res
}

// 思路：将所有的组合可能性列出并添加到哈希表中，
// 然后对字符串进行遍历，先判断两个字符的组合在哈希表中是否存在，存在则将值取出加到结果 res 中，并向后移2个字符；
// 不存在则将判断当前 1 个字符是否存在，存在则将值取出加到结果 res 中，并向后移 1 个字符。
// romanToInt2 returns int
func romanToInt2(s string) int {
	m := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var res int
	for i := 0; i < len(s); {
		if i+1 < len(s) && containsKey(m, string(s[i:i+2])) {
			res += m[string(s[i:i+2])]
			i += 2
		} else if containsKey(m, string(s[i:i+1])) {
			res += m[string(s[i:i+1])]
			i++
		}
	}

	return res
}

// containsKey return true if m contains k
func containsKey(m map[string]int, k string) bool {
	if _, ok := m[k]; ok {
		return true
	}
	return false
}
