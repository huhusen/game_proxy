package player

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

// player@game
type Game struct {
	cmd.Command
	Send2 struct {
		pkey string
	}
	Rec2 struct {
		Msg string `json:"msg"`
	}
}

func NewGame() *Game {
	u := Game{}
	u.Cmd = cmd.PlayerGame
	u.Zh = "????"
	return &u
}
func (c *Game) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
