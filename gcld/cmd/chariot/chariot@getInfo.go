package chariot

import (
	"encoding/hex"
	"fmt"
	"sockets-proxy/gcld"
	"sockets-proxy/gcld/cmd"
)

type GetInfo struct {
	Send struct {
		id string
		cp string
	}
}

// NewGetInfoRequest
//
//	@Description: 获取战车信息
//	@param id 战车id
//	@param cp ？
//	@return *ChariotGetInfo
func NewGetInfoRequest(id, cp string) *GetInfo {
	c := &GetInfo{}
	c.Send.cp = cp
	c.Send.id = id
	return c
}
func (l GetInfo) Data() []byte {
	body := fmt.Sprintf("id=%s&cp=%s", l.Send.id, l.Send.cp)
	return cmd.PacketData(cmd.ChariotGetBp, body)
}
func (l GetInfo) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := gcld.NewSendData(d)
	cmd.Print()
}
