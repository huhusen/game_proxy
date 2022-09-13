package gcld

import (
	"encoding/json"
	"fmt"
	"sockets-proxy/util"
	"sockets-proxy/util/binaryext"
)

type Data struct {
	Data     []byte `json:"-"`
	Len      uint32 `json:"len"`
	Command  string `json:"command"`
	Token    uint32 `json:"token"`
	Body     string `json:"body"`
	BodyByte []byte `json:"-"`
	Type     string `json:"-"`
}

func (r Data) String() string {
	body, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(body)
}
func (r Data) Bytes() []byte {
	body, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}
func NewRecvData(buf []byte) (r Data) {
	bytearray := binaryext.NewByteArray(buf)
	dataLen := bytearray.GetInt()
	if len(buf)-4 != int(dataLen) {
		return
	}
	command := bytearray.GetString(32)
	t := bytearray.GetInt()
	bodyBs := util.ZlibUnCompress(bytearray.Read(int(dataLen) - 36))
	body := string(bodyBs)
	r = Data{
		Len:      dataLen,
		Command:  command,
		Token:    t,
		Body:     body,
		BodyByte: bodyBs,
		Type:     "Recv",
	}
	return
}
func NewSendData(buf []byte) (s Data) {
	bytearray := binaryext.NewByteArray(buf)
	dataLen := bytearray.GetInt()
	if len(buf)-4 != int(dataLen) {
		return
	}
	command := bytearray.GetString(32)
	t := bytearray.GetInt()
	body := string(bytearray.Read(int(dataLen) - 36))
	s = Data{
		Data:    buf,
		Len:     dataLen,
		Command: command,
		Token:   t,
		Body:    body,
		Type:    "Send",
	}
	return
}
