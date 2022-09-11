package chariot

import (
	"encoding/hex"
	"fmt"
	"sockets-proxy/gcld/cmd"

	"sockets-proxy/gcld"
)

type ForgeBpInfo struct {
	Send struct {
		bpId string
	}
}

// NewForgeBpInfoRequest
//
//	@Description: 战车部件改造
//	@param bpId 部件id
//	@return *ChariotforgeBpInfo
func NewForgeBpInfoRequest(bpId string) *ForgeBpInfo {
	c := &ForgeBpInfo{}
	c.Send.bpId = bpId
	return c
}
func (l ForgeBpInfo) Data() []byte {
	body := fmt.Sprintf("bpId=%s", l.Send.bpId)
	return cmd.PacketData(cmd.ChariotforgeBp, body)
}
func (l ForgeBpInfo) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := gcld.NewSendData(d)
	cmd.Print()
}
