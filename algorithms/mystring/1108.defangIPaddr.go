package main

import (
	"fmt"
	"regexp"
)

func main() {
	address := "1.1.3.2"
	fmt.Println(defangIPaddr(address))
}

func defangIPaddr(address string) string {
	re := regexp.MustCompile(`\.`)
	return re.ReplaceAllString(address, "[.]")
}
