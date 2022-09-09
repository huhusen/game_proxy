package gcld

import (
	"encoding/hex"
	"fmt"
)

type ChariotGetInfo struct {
	Send struct {
		id string
		cp string
	}
}

//
// NewChariotGetInfoRequest
//  @Description: 获取战车信息
//  @param id 战车id
//  @param cp ？
//  @return *ChariotGetInfo
//
func NewChariotGetInfoRequest(id, cp string) *ChariotGetInfo {
	c := &ChariotGetInfo{}
	c.Send.cp = cp
	c.Send.id = id
	return c
}
func (l ChariotGetInfo) Data() []byte {
	body := fmt.Sprintf("id=%s&cp=%s", l.Send.id, l.Send.cp)
	return PacketData(ChariotGetBp, body)
}
func (l ChariotGetInfo) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := NewSendData(d)
	cmd.Print()
}
