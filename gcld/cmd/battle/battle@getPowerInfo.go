package battle

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

// battle@getPowerInfo
type GetPowerInfo struct {
	cmd.Command
	Send2 struct {
		string
	}
	Rec2 struct {
		TopLv          int    `json:"topLv"`
		NextPowerId    int    `json:"nextPowerId"`
		NextPowerName  string `json:"nextPowerName"`
		NextPowerIntro string `json:"nextPowerIntro"`
		Attackable     bool   `json:"attackable"`
		Type           int    `json:"type"`
		PowerId        int    `json:"powerId"`
		PowerName      string `json:"powerName"`
		PowerIntro     string `json:"powerIntro"`
		AutoCount      int    `json:"autoCount"`
		JxccPowerOpen  bool   `json:"jxccPowerOpen"`
		Npcs           []struct {
			AttLv      int           `json:"attLv"`
			NpcId      int           `json:"npcId"`
			NpcName    string        `json:"npcName"`
			NpcIntro   string        `json:"npcIntro"`
			Attackable bool          `json:"attackable"`
			Type       int           `json:"type"`
			Quality    int           `json:"quality"`
			FirstWin   int           `json:"firstWin"`
			FirstAtt   bool          `json:"firstAtt"`
			Terrain    int           `json:"terrain"`
			BatInfos   []interface{} `json:"batInfos,omitempty"`
			Attacked   bool          `json:"attacked"`
		} `json:"npcs"`
		RewardNpcsPos []struct {
		} `json:"rewardNpcsPos"`
	}
}

func NewGetPowerInfo() *GetPowerInfo {
	u := GetPowerInfo{}
	u.Cmd = cmd.BattleGetPowerInfo
	u.Zh = "????"
	return &u
}
func (c *GetPowerInfo) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
