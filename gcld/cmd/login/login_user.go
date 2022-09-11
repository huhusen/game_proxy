package login

import (
	"fmt"
	"sockets-proxy/gcld/cmd"

	"sockets-proxy/util"
)

type UserInfo struct {
	Send struct {
		userkey string
	}
	Rec struct {
		sessionId string
	}
}

func NewUserInfoRequest(userkey string) *UserInfo {
	return &UserInfo{Send: struct{ userkey string }{userkey: userkey}}
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}
func (l UserInfo) UpdateRec(data []byte) {
	m := util.NewJSon(data)
	l.Rec.sessionId = m.Get("sessionId").MustString("")
}
func (l UserInfo) UpdateSend(data string) {
	m := util.NewMap(data)
	l.Send.userkey = m["userkey"]
}
func (l UserInfo) Data() []byte {
	body := fmt.Sprintf("userkey=%s", l.Send.userkey)
	return cmd.PacketData(cmd.LoginUser, body)
}
