package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func IP2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func main() {
	s := "127.0.0.1"
	fmt.Println(IP2long(s))
}
