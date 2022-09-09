package gcld

import (
	"encoding/hex"
	"fmt"
)

type ChariotGetBpInfo struct {
	Send struct {
		bpId string
		id   string
	}
}

//
// NewChariotGetBpInfoRequest
//  @Description:获取战车部件信息
//  @param bpId 部件id
//  @param id 战车id
//  @return *ChariotGetBpInfo
//
func NewChariotGetBpInfoRequest(bpId, id string) *ChariotGetBpInfo {
	c := &ChariotGetBpInfo{}
	c.Send.bpId = bpId
	c.Send.id = id
	return c
}
func (l ChariotGetBpInfo) Data() []byte {
	body := fmt.Sprintf("id=%s&bpId=%s", l.Send.id, l.Send.bpId)
	return PacketData(ChariotGetBp, body)
}
func (l ChariotGetBpInfo) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := NewSendData(d)
	cmd.Print()
}
