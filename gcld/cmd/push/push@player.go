package push

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

type Player struct {
	cmd.Command
	Rec2 struct {
		Update struct {
			HasFirstPayPkg bool `json:"hasFirstPayPkg"`
		} `json:"update"`
	}
}

func NewPlayer() *Player {
	u := Player{}
	u.Cmd = cmd.PushPlayer
	u.Zh = "【推送】用户信息"
	//u.Send = u.Send2
	u.Rec = u.Rec2
	return &u
}
func (c *Player) Update() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
