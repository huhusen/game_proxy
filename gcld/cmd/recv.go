package cmd

import (
	"encoding/json"
	"fmt"
	"sockets-proxy/util"
	"sockets-proxy/util/binaryext"
)

type RecData struct {
	Data     []byte `json:"-"`
	Len      uint32 `json:"len"`
	Command  CMD    `json:"command"`
	Token    uint32 `json:"token"`
	Body     string `json:"body"`
	BodyByte []byte `json:"-"`
}

func (r RecData) String() string {
	body, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(body)
}
func (r RecData) Bytes() []byte {
	body, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}
func NewRecvData(buf []byte) (r RecData) {
	bytearray := binaryext.NewByteArray(buf)
	dataLen := bytearray.GetInt()
	if len(buf)-4 != int(dataLen) {
		return
	}
	command := bytearray.GetString(32)
	t := bytearray.GetInt()
	bodyBs := util.ZlibUnCompress(bytearray.Read(int(dataLen) - 36))
	body := string(bodyBs)
	r = RecData{
		Len:      dataLen,
		Command:  CMD(command),
		Token:    t,
		Body:     body,
		BodyByte: bodyBs,
	}
	return
}

//func (r RecData) Print() {
//	if r.Len == 0 {
//		return
//	}
//	log.Log.Printf("RecvData:【%s】\nData:%s\nLen:%d\nCommand:%s\nToken:%d\nBody:%s \n", ToCMD(r.Command),
//		hex.EncodeToString(r.Data), r.Len, r.Command, r.Token, r.Body)
//}
