package building

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

type GetBuildingInfo struct {
	cmd.Command
	Send2 struct {
		types string `json:"type"`
	}
	Rec2 struct {
		Luban        int `json:"luban"`
		SlaveNum     int `json:"slaveNum"`
		BuyGrabCost  int `json:"buyGrabCost"`
		BuyGrabNum   int `json:"buyGrabNum"`
		BuyGrabState int `json:"buyGrabState"`
		Progress     int `json:"progress"`
		BuyCost      int `json:"buyCost"`
		Buildings    []struct {
			Status     int    `json:"status,omitempty"`
			Harvest    int    `json:"harvest,omitempty"`
			CanChanged bool   `json:"canChanged,omitempty"`
			IsChanged  bool   `json:"isChanged,omitempty"`
			Type       int    `json:"type"`
			ID         int    `json:"id"`
			Name       string `json:"name"`
			SpeedUpNum int    `json:"speedUpNum,omitempty"`
			TotalUpNum int    `json:"totalUpNum,omitempty"`
			Lv         int    `json:"lv,omitempty"`
			Intro      string `json:"intro,omitempty"`
			Pos        int    `json:"pos"`
			OutputType int    `json:"outputType,omitempty"`
			ResType    int    `json:"resType,omitempty"`
			Output     int    `json:"output,omitempty"`
			IsNew      bool   `json:"isNew"`
			HasEvent   bool   `json:"hasEvent,omitempty"`
			Upgrade    struct {
				UpgradeEnable bool `json:"upgradeEnable"`
				Time          int  `json:"time"`
				Cost          []struct {
					Type  int `json:"type"`
					Value int `json:"value"`
				} `json:"cost"`
			} `json:"upgrade,omitempty"`
		} `json:"buildings"`
		Tip         bool `json:"tip"`
		TotalOutput []struct {
			Type   int `json:"type"`
			Output int `json:"output"`
		} `json:"totalOutput"`
		AutoUpbuilding struct {
			Times  int `json:"times"`
			AreaID int `json:"areaId"`
			State  int `json:"state"`
		} `json:"autoUpbuilding"`
		AdditionMode  int `json:"additionMode"`
		AdditionRate  int `json:"additionRate"`
		AdditionCd    int `json:"additionCd"`
		BuildingWorks []struct {
			WorkID int `json:"workId"`
			State  int `json:"state"`
		} `json:"buildingWorks"`
		FreeConsNum int `json:"freeConsNum"`
		HasBandit   int `json:"hasBandit"`
	}
}

func NewGetBuildingInfo() *GetBuildingInfo {
	u := GetBuildingInfo{}
	u.Cmd = cmd.BuildingGetBuildingInfo
	u.Zh = "【建筑】获取建筑信息"
	return &u
}

func (c *GetBuildingInfo) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
