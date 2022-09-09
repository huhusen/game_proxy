package gcld

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sockets-proxy/binaryext"
)

type sendData struct {
	Data    []byte `json:"-"`
	Len     uint32 `json:"len"`
	Command CMD    `json:"command"`
	Token   uint32 `json:"token"`
	Body    string `json:"body"`
}

func (s sendData) String() string {
	body, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(body)
}
func (s sendData) Bytes() []byte {
	body, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}
func NewSendData(buf []byte) (s sendData) {
	bytearray := binaryext.NewByteArray(buf)
	dataLen := bytearray.GetInt()
	if len(buf)-4 != int(dataLen) {
		return
	}
	cmd := bytearray.GetString(32)
	t := bytearray.GetInt()
	body := string(bytearray.Read(int(dataLen) - 36))
	s = sendData{
		Data:    buf,
		Len:     dataLen,
		Command: CMD(cmd),
		Token:   t,
		Body:    body,
	}
	return
}
func (s sendData) Print() {
	if s.Len == 0 {
		return
	}
	fmt.Printf("SendData:【%s】\nData:%s\nLen:%d\nCommand:%s\nToken:%d\nBody:%s \n", ToCMD(s.Command),
		hex.EncodeToString(s.Data), s.Len, s.Command, s.Token, s.Body)
}
