package gcld

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sockets-proxy/binaryext"
	"sockets-proxy/util"
)

type recData struct {
	Data    []byte `json:"-"`
	Len     uint32 `json:"len"`
	Command CMD    `json:"command"`
	Token   uint32 `json:"token"`
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
	bytearray := binaryext.NewByteArray(buf)
	dataLen := bytearray.GetInt()
	if len(buf)-4 != int(dataLen) {
		return
	}
	cmd := bytearray.GetString(32)
	t := bytearray.GetInt()
	body := string(util.ZlibUnCompress(bytearray.Read(int(dataLen) - 36)))
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
	fmt.Printf("RecvData:【%s】\nData:%s\nLen:%d\nCommand:%s\nToken:%d\nBody:%s \n", ToCMD(r.Command),
		hex.EncodeToString(r.Data), r.Len, r.Command, r.Token, r.Body)
}
