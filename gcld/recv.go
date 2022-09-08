package gcld

import (
	"encoding/json"
	"fmt"
	"sockets-proxy/log"
	"sockets-proxy/util"
)

type recData struct {
	Len     int32  `json:"len"`
	Command CMD    `json:"command"`
	Token   int32  `json:"token"`
	Body    string `json:"body"`
}

func (r recData) String() string {
	body, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(body)
}
func (r recData) Bytes() []byte {
	body, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}
func NewRecvData(buf []byte) (r recData) {
	reader := util.NewBinaryReader(buf)
	dataLen := reader.ReadInt32()
	if len(buf)-4 != int(dataLen) {
		return
	}
	cmd := string(reader.ReadBytes(32))
	t := reader.ReadInt32()
	body := string(util.ZlibUnCompress(reader.ReadBytes(int(dataLen) - 36)))
	r = recData{
		Len:     dataLen,
		Command: CMD(cmd),
		Token:   t,
		Body:    body,
	}
	return
}
func (r recData) Print() {
	if r.Len == 0 {
		return
	}
	log.Log.Printf("\nRecvData:\nLen:%d\nCommand:%s\nToken:%d\nBody:%s \n", r.Len, r.Command, r.Token, r.Body)
}
