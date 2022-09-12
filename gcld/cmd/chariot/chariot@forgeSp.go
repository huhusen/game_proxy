package chariot

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

//  ForgeSp
//  @Description:  战车升级

type ForgeSp struct {
	cmd.Command
	Send2 struct {
		zcjd   string
		bpId   string
		spId   string
		buffId string
	}
	Rec2 struct {
		defUp        string
		meteoriteNum string
		blueprintNum string
		extrabpnum   string
		bufflist     []string
	}
}

func NewForgeSp() *ForgeSp {
	p := &ForgeSp{}
	p.Zh = "【战车】战车升级"
	p.Cmd = cmd.ChariotForgeSp
	p.Rec = p.Rec2
	return p
}

// NewForgeSpRequest
//
//	@Description: 战车升级
//	@param bpId 部件id
//	@param spId 火力id
//	@return *ForgeSp
func NewForgeSpRequest(bpId, spId string) *ForgeSp {
	c := &ForgeSp{}
	c.Send2.bpId = bpId
	c.Send2.spId = spId
	return c
}
func (l *ForgeSp) Data() []byte {
	body := fmt.Sprintf("zcjd=0&bpId=%s&spId=%s&buffId=0", l.Send2.bpId, l.Send2.spId)
	return l.PacketData(body)
}
func (c *ForgeSp) Update() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
