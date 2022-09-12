package chariot

import (
	"fmt"
	"sockets-proxy/gcld/cmd"
)

type ForgeBp struct {
	cmd.Command
	Send2 struct {
		bpId string
	}
}

func NewForgeBpInfo() *ForgeBp {
	u := ForgeBp{}
	u.Cmd = cmd.ChariotforgeBp
	u.Zh = "【战车】战车部件改造"
	return &u
}

// NewForgeBpInfoRequest
//
//	@Description: 战车部件改造
//	@param bpId 部件id
//	@return *ChariotforgeBpInfo
func NewForgeBpInfoRequest(bpId string) *ForgeBp {
	c := &ForgeBp{}
	c.Send2.bpId = bpId
	return c
}
func (l *ForgeBp) Data() []byte {
	body := fmt.Sprintf("bpId=%s", l.Send2.bpId)
	return l.PacketData(body)
}
