package login

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
	"sockets-proxy/util"
)

type UserInfo struct {
	cmd.Command
	Send2 struct {
		userkey string
	}
	Rec2 struct {
		SessionID string `json:"sessionId"`
	}
}

func NewUserInfoRequest(userkey string) *UserInfo {
	return &UserInfo{Send2: struct{ userkey string }{userkey: userkey}}
}

func NewUserInfo() *UserInfo {
	u := UserInfo{}
	u.Cmd = cmd.LoginUser
	u.Zh = "【登录】用户登录"
	u.Send = u.Send2
	u.Rec = u.Rec2
	return &u
}

func (c *UserInfo) Data() []byte {
	body := fmt.Sprintf("userkey=%s", c.Send2.userkey)
	return c.PacketData(body)
}

func (c *UserInfo) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
func (c *UserInfo) Update2() {
	util.Map2Struct(c.Send.(string), &c.Send2)
	fmt.Println()
}
