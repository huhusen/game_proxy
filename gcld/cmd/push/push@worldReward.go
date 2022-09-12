package push

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

// push@worldReward
type WorldReward struct {
	cmd.Command
	Send2 struct {
	}
	Rec2 struct {
		WorldReward []struct {
			Quality  int `json:"quality"`
			LeftTime int `json:"leftTime"`
		} `json:"worldReward"`
	}
}

func NewWorldReward() *WorldReward {
	u := WorldReward{}
	u.Cmd = cmd.PushWorldReward
	u.Zh = "????"
	return &u
}
func (c *WorldReward) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
