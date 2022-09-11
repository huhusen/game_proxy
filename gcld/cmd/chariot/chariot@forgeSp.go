package chariot

import (
	"encoding/hex"
	"fmt"
	"sockets-proxy/gcld/cmd"

	"sockets-proxy/gcld"
)

//  ForgeSpInfo
//  @Description:  战车升级

type ForgeSpInfo struct {
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

// NewForgeSpInfoRequest
//
//	@Description: 战车升级
//	@param bpId 部件id
//	@param spId 火力id
//	@return *ForgeSpInfo
func NewForgeSpInfoRequest(bpId, spId string) *ForgeSpInfo {
	c := &ForgeSpInfo{}
	c.Send.bpId = bpId
	c.Send.spId = spId
	return c
}
func (l ForgeSpInfo) Data() []byte {
	body := fmt.Sprintf("zcjd=0&bpId=%s&spId=%s&buffId=0", l.Send.bpId, l.Send.spId)
	return cmd.PacketData(cmd.ChariotForgeSp, body)
}
func (l ForgeSpInfo) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := gcld.NewSendData(d)
	cmd.Print()
}
