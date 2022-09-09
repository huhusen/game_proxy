package gcld

import (
	"encoding/hex"
	"fmt"
)

//  ChariotForgeSpInfo
//  @Description:  战车升级

type ChariotForgeSpInfo struct {
	Send struct {
		zcjd   string
		bpId   string
		spId   string
		buffId string
	}
	Rec struct {
		defUp        string
		meteoriteNum string
		blueprintNum string
		extrabpnum   string
		bufflist     []string
	}
}

//
// NewChariotForgeSpInfoRequest
//  @Description: 战车升级
//  @param bpId 部件id
//  @param spId 火力id
//  @return *ChariotForgeSpInfo
//
func NewChariotForgeSpInfoRequest(bpId, spId string) *ChariotForgeSpInfo {
	c := &ChariotForgeSpInfo{}
	c.Send.bpId = bpId
	c.Send.spId = spId
	return c
}
func (l ChariotForgeSpInfo) Data() []byte {
	body := fmt.Sprintf("zcjd=0&bpId=%s&spId=%s&buffId=0", l.Send.bpId, l.Send.spId)
	return PacketData(ChariotForgeSp, body)
}
func (l ChariotForgeSpInfo) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := NewSendData(d)
	cmd.Print()
}
