package chariot

import (
	"encoding/hex"
	"fmt"
	"sockets-proxy/gcld/cmd"

	"sockets-proxy/gcld"
)

type GetBpInfo struct {
	Send struct {
		bpId string
		id   string
	}
}

// NewGetBpInfoRequest
//
//	@Description:获取战车部件信息
//	@param bpId 部件id
//	@param id 战车id
//	@return *GetBpInfo
func NewGetBpInfoRequest(bpId, id string) *GetBpInfo {
	c := &GetBpInfo{}
	c.Send.bpId = bpId
	c.Send.id = id
	return c
}
func (l GetBpInfo) Data() []byte {
	body := fmt.Sprintf("id=%s&bpId=%s", l.Send.id, l.Send.bpId)
	return cmd.PacketData(cmd.ChariotGetBp, body)
}
func (l GetBpInfo) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := gcld.NewSendData(d)
	cmd.Print()
}
