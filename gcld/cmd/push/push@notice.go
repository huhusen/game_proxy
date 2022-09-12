package push

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

type Notice struct {
	cmd.Command
	Rec2 struct {
		Notice struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"notice"`
	}
}

func NewNotice() *Notice {
	u := Notice{}
	u.Cmd = cmd.PushNotice
	u.Zh = "【推送】通知"
	//u.Send = u.Send2
	u.Rec = u.Rec2
	return &u
}
func (c *Notice) Update() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
