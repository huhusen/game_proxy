package push

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

// push@rightNotice
type RightNotice struct {
	cmd.Command
	Send2 struct {
	}
	Rec2 struct {
		RightNotice []struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"rightNotice"`
	}
}

func NewRightNotice() *RightNotice {
	u := RightNotice{}
	u.Cmd = cmd.PushRightNotice
	u.Zh = "????"
	return &u
}
func (c *RightNotice) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
