package mystring

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	slices := strings.Split(IP, ".")
	if len(slices) == 4 && isIPv4(slices) == true {
		return "IPv4"
	}

	lowerIP := strings.ToLower(IP)
	parts := strings.Split(lowerIP, ":")
	if len(parts) == 8 && isIPv6(parts) == true {
		return "IPv6"
	}

	return "Neither"
}

func isIPv4(slices []string) bool {
	for _, num := range slices {
		// 过滤非int值
		numInt, err := strconv.Atoi(num)
		if err != nil {
			return false
		}
		if numInt < 0 || numInt > 255 {
			return false
		}
		// 过滤0开头的段，如: 01.1.1.1
		if strconv.Itoa(numInt) != num {
			return false
		}
	}
	return true
}

func isIPv6(parts []string) bool {
	for _, part := range parts {
		if len(part) == 0 || len(part) > 4 {
			return false
		}
		for _, v := range part {
			if v < '0' || (v > '9' && v < 'a') || v > 'f' {
				return false
			}
		}
	}
	return true
}
