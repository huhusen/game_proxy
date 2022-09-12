package push

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

// push@chat
type Chat struct {
	cmd.Command
	Send2 struct {
	}
	Rec2 struct {
		ChatSend []struct {
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
	u.Zh = "????"
	return &u
}
func (c *Chat) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
