package chariot

import (
	"fmt"
	"sockets-proxy/gcld/cmd"
)

type GetInfo struct {
	cmd.Command
	Send2 struct {
		id string
		cp string
	}
}

func NewGetInfo() *GetInfo {
	p := &GetInfo{}
	p.Zh = "【战车】获取战车信息"
	p.Cmd = cmd.ChariotGetBpInfo
	//p.Rec = p.Rec2
	return p
}

// NewGetInfoRequest
//
//	@Description: 获取战车信息
//	@param id 战车id
//	@param cp ？
//	@return *ChariotGetInfo
func NewGetInfoRequest(id, cp string) *GetInfo {
	c := &GetInfo{}
	c.Send2.cp = cp
	c.Send2.id = id
	return c
}
func (l GetInfo) Data() []byte {
	body := fmt.Sprintf("id=%s&cp=%s", l.Send2.id, l.Send2.cp)
	return l.PacketData(body)
}
