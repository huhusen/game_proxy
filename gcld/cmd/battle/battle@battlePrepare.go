package battle

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

// battle@battlePrepare
type BattlePrepare struct {
	cmd.Command
	Send2 struct {
		generalId   string
		targetId    string
		terrainType string
		types       string
	}
	Rec2 struct {
		Side             int    `json:"side"`
		BattleType       int    `json:"battleType"`
		Terrain          int    `json:"terrain"`
		TerrainPic       int    `json:"terrainPic"`
		TargetId         int    `json:"targetId"`
		TargetName       string `json:"targetName"`
		FreePhantomCount int    `json:"freePhantomCount"`
		Name             string `json:"name"`
		Intro            string `json:"intro"`
		Bat              int    `json:"bat"`
		Food             int    `json:"food"`
		Color            bool   `json:"color"`
		PlayerId         int    `json:"playerId"`
		PlayerName       string `json:"playerName"`
		PlayerPic        int    `json:"playerPic"`
		PlayerLv         int    `json:"playerLv"`
		NpcId            int    `json:"npcId"`
		NpcName          string `json:"npcName"`
		NpcPic           string `json:"npcPic"`
		NpcLv            int    `json:"npcLv"`
		AttGenerals      []struct {
			Index           int    `json:"index"`
			GeneralId       int    `json:"generalId"`
			GeneralName     string `json:"generalName"`
			GeneralLv       int    `json:"generalLv"`
			HelperGeneralId int    `json:"helperGeneralId"`
			Quality         int    `json:"quality"`
			TroopId         int    `json:"troopId"`
			TroopType       int    `json:"troopType"`
			State           int    `json:"state"`
			GeneralPic      string `json:"generalPic"`
			ArmyHp          int    `json:"armyHp"`
			ArmyHpMax       int    `json:"armyHpMax"`
			NeedTime        int    `json:"needTime"`
		} `json:"attGenerals"`
		DefGenerals []struct {
			GeneralId   int    `json:"generalId"`
			GeneralName string `json:"generalName"`
			Att         int    `json:"att"`
			GeneralLv   int    `json:"generalLv"`
			TroopId     int    `json:"troopId"`
			TroopType   int    `json:"troopType"`
			GeneralPic  string `json:"generalPic"`
			Quality     int    `json:"quality"`
			ArmyHp      int    `json:"armyHp"`
			ArmyHpMax   int    `json:"armyHpMax"`
		} `json:"defGenerals"`
		Cost1  int `json:"cost1"`
		Cost2  int `json:"cost2"`
		Time   int `json:"time"`
		Mode   int `json:"mode"`
		AutoSt int `json:"autoSt"`
	}
}

func NewBattlePrepare() *BattlePrepare {
	u := BattlePrepare{}
	u.Cmd = cmd.BattleBattlePrepare
	u.Zh = "????"
	return &u
}
func (c *BattlePrepare) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
