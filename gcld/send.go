package gcld

import (
	"encoding/json"
	"fmt"
	"sockets-proxy/log"
	"sockets-proxy/util"
)

type sendData struct {
	Len     int32  `json:"len"`
	Command CMD    `json:"command"`
	Token   int32  `json:"token"`
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
	reader := util.NewBinaryReader(buf)
	dataLen := reader.ReadInt32()
	if len(buf)-4 != int(dataLen) {
		return
	}
	cmd := string(reader.ReadBytes(32))
	t := reader.ReadInt32()
	body := string(reader.ReadBytes(int(dataLen) - 36))
	s = sendData{
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
	log.Log.Printf("\nSendData:\nLen:%d\nCommand:%s\nToken:%d\nBody:%s \n", s.Len, s.Command, s.Token, s.Body)
}
