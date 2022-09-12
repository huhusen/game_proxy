package push

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

type Chat struct {
	cmd.Command
	Rec2 struct {
		ChatSend struct {
			IsGm  bool   `json:"isGm"`
			From  string `json:"from"`
			To    string `json:"to"`
			Type  string `json:"type"`
			Msg   string `json:"msg"`
			Voice bool   `json:"voice"`
		} `json:"chatSend"`
	}
}

func NewChat() *Chat {
	u := Chat{}
	u.Cmd = cmd.PushChat
	u.Zh = "【推送】聊天"
	//u.Send = u.Send2
	u.Rec = u.Rec2
	return &u
}

func (c *Chat) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}

//func (c *GetPlayerList) Update2() {
//	util.Map2Struct(c.Send.(string), &c.Send2)
//	fmt.Println()
//}
