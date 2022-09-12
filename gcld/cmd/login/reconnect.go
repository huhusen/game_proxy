package login

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

type Reconnect struct {
	cmd.Command
	Send2 struct {
		sessionId string
	}
	Rec2 struct {
		Msg string `json:"msg"`
	}
}

func NewReconnect() *Reconnect {
	u := Reconnect{}
	u.Cmd = cmd.Reconnect
	u.Zh = "????"
	return &u
}
func (c *Reconnect) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
