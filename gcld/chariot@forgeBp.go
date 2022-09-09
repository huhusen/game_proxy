package gcld

import (
	"encoding/hex"
	"fmt"
)

type ChariotforgeBpInfo struct {
	Send struct {
		bpId string
	}
}

//
// NewChariotforgeBpInfoRequest
//  @Description: 战车部件改造
//  @param bpId 部件id
//  @return *ChariotforgeBpInfo
//
func NewChariotforgeBpInfoRequest(bpId string) *ChariotforgeBpInfo {
	c := &ChariotforgeBpInfo{}
	c.Send.bpId = bpId
	return c
}
func (l ChariotforgeBpInfo) Data() []byte {
	body := fmt.Sprintf("bpId=%s", l.Send.bpId)
	return PacketData(ChariotforgeBp, body)
}
func (l ChariotforgeBpInfo) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := NewSendData(d)
	cmd.Print()
}
