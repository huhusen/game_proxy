package util

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"net"
	"strings"
)

func NewBuf(len int) []byte {
	return make([]byte, len)
}

//dynamic := make(map[string]interface{})

func NewMap(param interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	switch param.(type) {
	case string:
		str := param.(string)
		arr := strings.Split(str, "=")
		if len(arr)%2 != 0 {
			return nil
		} else {
			for i := 0; i < len(arr); i = i + 2 {
				m[arr[i]] = arr[i+1]
			}
		}
	case []byte:
		json.Unmarshal(param.([]byte), &m)
		m = m["action"].(map[string]interface{})
		m = m["data"].(map[string]interface{})
	}
	return m
}

// ZlibCompress 进行zlib压缩
func ZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

// ZlibUnCompress 进行zlib解压缩
func ZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

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
