package main

import "fmt"

func reverseString(s []byte) string {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
	return string(s)
}

func main() {
	s := "hello"
	fmt.Println(reverseString([]byte(s)))
}
