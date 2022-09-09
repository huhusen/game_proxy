package gcld

import (
	"fmt"
	"sockets-proxy/util"
)

type LoginUserInfo struct {
	Send struct {
		userkey string
	}
	Rec struct {
		sessionId string
	}
}

func NewLoginUserInfoRequest(userkey string) *LoginUserInfo {
	return &LoginUserInfo{Send: struct{ userkey string }{userkey: userkey}}
}

func NewLoginUserInfo() *LoginUserInfo {
	return &LoginUserInfo{}
}
func (l LoginUserInfo) UpdateRec(data []byte) {
	m := util.NewMap(data)
	l.Rec.sessionId = m["sessionId"].(string)
}
func (l LoginUserInfo) UpdateSend(data string) {
	m := util.NewMap(data)
	l.Send.userkey = m["userkey"].(string)
}
func (l LoginUserInfo) Data() []byte {
	body := fmt.Sprintf("userkey=%s", l.Send.userkey)
	return PacketData(LoginUser, body)
}
