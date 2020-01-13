package mystring

import "regexp"

func defangIPaddr(address string) string {
	re := regexp.MustCompile(`\.`)
	return re.ReplaceAllString(address, "[.]")
}
