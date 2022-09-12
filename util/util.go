package util

import (
	"bytes"
	"compress/zlib"
	"github.com/bitly/go-simplejson"
	"github.com/goinggo/mapstructure"
	"io"
	"net"
	"strings"
)

func NewBuf(len int) []byte {
	return make([]byte, len)
}

//dynamic := make(map[string]interface{})

func NewMap(param string) map[string]string {
	m := make(map[string]string)
	arr := strings.Split(param, "=")
	if len(arr)%2 != 0 {
		return nil
	} else {
		for i := 0; i < len(arr); i = i + 2 {
			if arr[i] != "type" {
				m[arr[i]] = arr[i+1]
			} else {
				m["types"] = arr[i+1]
			}
		}
	}
	return m
}
func NewJSon(data []byte) *simplejson.Json {
	js, err := simplejson.NewJson(data)
	if err != nil {
		return nil
	}
	return js.Get("action").Get("data")
}
func Byte2Struct(data []byte, s interface{}) {
	j := NewJSon(data)
	m, _ := j.Map()
	mapstructure.Decode(m, s)

}
func Map2Struct(param string, s interface{}) {
	mapstructure.Decode(NewMap(param), &s)
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
