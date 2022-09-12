package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sockets-proxy/util/binaryext"
)

type SendData struct {
	Data    []byte `json:"-"`
	Len     uint32 `json:"len"`
	Command CMD    `json:"command"`
	Token   uint32 `json:"token"`
	Body    string `json:"body"`
}

func NewSendData(buf []byte) (s SendData) {
	bytearray := binaryext.NewByteArray(buf)
	dataLen := bytearray.GetInt()
	if len(buf)-4 != int(dataLen) {
		return
	}
	command := bytearray.GetString(32)
	t := bytearray.GetInt()
	body := string(bytearray.Read(int(dataLen) - 36))
	s = SendData{
		Data:    buf,
		Len:     dataLen,
		Command: CMD(command),
		Token:   t,
		Body:    body,
	}
	return
}
func (s SendData) String() string {
	body, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(body)
}
func (s SendData) Bytes() []byte {
	body, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}

func (s SendData) Print(name string) {
	if s.Len == 0 {
		return
	}
	fmt.Printf("SendData:【%s】\nData:%s\nLen:%d\nCommand:%s\nToken:%d\nBody:%s \n", name,
		hex.EncodeToString(s.Data), s.Len, s.Command, s.Token, s.Body)
}
