package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "09cbAa12d98xYz~#"
	fmt.Println(sortStringOnly(str))
}

func sortStringOnly(str string) string {
	var (
		letter, other []string
		flags         []bool
	)
	for _, char := range []byte(str) {
		// a-z || A-Z
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
