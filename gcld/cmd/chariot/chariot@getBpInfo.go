package chariot

import (
	"fmt"
	"sockets-proxy/gcld/cmd"
)

type GetBpInfo struct {
	cmd.Command
	Send2 struct {
		bpId string
		id   string
	}
}

func NewGetBpInfo() *GetBpInfo {
	p := &GetBpInfo{}
	p.Zh = "【战车】获取战车部件信息"
	p.Cmd = cmd.ChariotGetBpInfo
	//p.Rec = p.Rec2
	return p
}

// NewGetBpInfoRequest
//
//	@Description:获取战车部件信息
//	@param bpId 部件id
//	@param id 战车id
//	@return *GetBpInfo
func NewGetBpInfoRequest(bpId, id string) *GetBpInfo {
	c := &GetBpInfo{}
	c.Send2.bpId = bpId
	c.Send2.id = id
	return c
}
func (l *GetBpInfo) Data() []byte {
	body := fmt.Sprintf("id=%s&bpId=%s", l.Send2.id, l.Send2.bpId)
	return l.PacketData(body)

}
