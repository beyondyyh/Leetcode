package tree

import "strings"

func isValidSerialization(preorder string) bool {
	var leaves, node int
	pres := strings.Split(preorder, ",")
	for i, s := range pres {
		if s == "#" {
			leaves++
		} else {
			node++
		}
		if leaves > node+1 {
			return false
		}
		if leaves == node+1 && i < len(pres)-1 {
			return false
		}
	}

	if leaves == node+1 {
		return true
	}
	return false
}
