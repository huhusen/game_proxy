package gcld

import (
	"fmt"
	"sockets-proxy/util"
)

type RecvData struct {
	Len     int32
	Command string
	Token   int32
	Body    string
}

func NewRecvData(buf []byte) (r RecvData) {
	reader := util.NewBinaryReader(buf)
	dataLen := reader.ReadInt32()
	if len(buf)-4 != int(dataLen) {
		return
	}
	cmd := string(reader.ReadBytes(32))
	t := reader.ReadInt32()
	body := string(util.ZlibUnCompress(reader.ReadBytes(int(dataLen) - 36)))
	r = RecvData{
		Len:     dataLen,
		Command: cmd,
		Token:   t,
		Body:    body,
	}
	return
}
func (r RecvData) Print() {
	if r.Len == 0 {
		return
	}
	fmt.Printf("\nRecvData:\nLen:%d\nCommand:%s\nToken:%d\nBody:%s \n", r.Len, r.Command, r.Token, r.Body)
}
