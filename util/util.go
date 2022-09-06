package util

import (
	"net"
	"strings"
)

// ByteToInt 字节转整型
func ByteToInt(input []byte) int32 {
	return int32(input[0]&0xFF)<<8 | int32(input[1]&0xFF)
}
func IsIpV4(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return ip != nil && strings.Contains(ipAddr, ".")
}

func IsIpV6(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return ip != nil && strings.Contains(ipAddr, ":")
}
