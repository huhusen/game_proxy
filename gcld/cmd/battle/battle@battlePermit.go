package battle

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

// battle@battlePermit
type BattlePermit struct {
	cmd.Command
	Send2 struct {
		generalId string
		targetId  string
		reserve   string
		types     string
	}
	Rec2 struct {
		Battle bool `json:"battle"`
		Side   int  `json:"side"`
	}
}

func NewBattlePermit() *BattlePermit {
	u := BattlePermit{}
	u.Cmd = cmd.BattleBattlePermit
	u.Zh = "????"
	return &u
}
func (c *BattlePermit) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
