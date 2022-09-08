package gcld

import (
	"fmt"
	"sockets-proxy/util"
)

type SendData struct {
	Len     int32
	Command string
	Token   int32
	Body    string
}

func NewSendData(buf []byte) (s SendData) {
	reader := util.NewBinaryReader(buf)
	dataLen := reader.ReadInt32()
	if len(buf)-4 != int(dataLen) {
		return
	}
	cmd := string(reader.ReadBytes(32))
	t := reader.ReadInt32()
	body := string(reader.ReadBytes(int(dataLen) - 36))
	s = SendData{
		Len:     dataLen,
		Command: cmd,
		Token:   t,
		Body:    body,
	}
	return
}
func (s SendData) Print() {
	if s.Len == 0 {
		return
	}
	fmt.Printf("\nSendData:\nLen:%d\nCommand:%s\nToken:%d\nBody:%s \n", s.Len, s.Command, s.Token, s.Body)
}
