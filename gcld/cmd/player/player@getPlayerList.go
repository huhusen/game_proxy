package player

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

type GetPlayerList struct {
	cmd.Command
	Rec2 struct {
		MaxRoleNum int `json:"maxRoleNum"`
		PlayerList []struct {
			PlayerID      int    `json:"playerId"`
			PlayerLv      int    `json:"playerLv"`
			PlayerName    string `json:"playerName"`
			LastLoginTime int64  `json:"lastLoginTime"`
			Pic           int    `json:"pic"`
			ConsumLv      int    `json:"consumLv"`
			IsDelete      bool   `json:"isDelete"`
			ForceID       int    `json:"forceId"`
			DefaultPay    int    `json:"defaultPay"`
			DelegateCd    int    `json:"delegateCd"`
		} `json:"playerList"`
		WeiName    string `json:"weiName"`
		ShuName    string `json:"shuName"`
		WuName     string `json:"wuName"`
		SingleRole bool   `json:"singleRole"`
		MaxRole    bool   `json:"maxRole"`
		ServerName string `json:"serverName"`
		ServerID   string `json:"serverId"`
	}
}

func NewGetPlayerList() *GetPlayerList {
	p := &GetPlayerList{}
	p.Zh = "【玩家】获取用户列表信息"
	p.Cmd = cmd.PlayerGetPlayerList
	p.Rec = p.Rec2
	return p
}

func (c *GetPlayerList) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}

//func (c *GetPlayerList) Update2() {
//	util.Map2Struct(c.Send.(string), &c.Send2)
//	fmt.Println()
//}
