package building

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
	"sockets-proxy/util"
)

type GetMainCityInfo struct {
	cmd.Command
	Send2 struct {
		param string
	}
	Rec2 struct {
		Areas []struct {
			ID           int  `json:"id"`
			IsOpen       bool `json:"isOpen"`
			HasEvent     bool `json:"hasEvent,omitempty"`
			AdditionMode int  `json:"additionMode,omitempty"`
			AdditionRate int  `json:"additionRate,omitempty"`
			AdditionCd   int  `json:"additionCd,omitempty"`
			PbwNum       int  `json:"pbwNum,omitempty"`
			IsBuilding   int  `json:"isBuilding,omitempty"`
			CanReform    int  `json:"canReform,omitempty"`
			Status       int  `json:"status,omitempty"`
			TotalOutput  []struct {
				Type   int `json:"type"`
				Output int `json:"output"`
			} `json:"totalOutput,omitempty"`
			SilkAdditionMode int `json:"silkAdditionMode,omitempty"`
			SilkAdditionRate int `json:"silkAdditionRate,omitempty"`
			SilkAdditionCd   int `json:"silkAdditionCd,omitempty"`
			TotalSilkOutPut  int `json:"totalSilkOutPut,omitempty"`
			Addition         int `json:"addition,omitempty"`
			OfficerOutput    []struct {
				Type  int `json:"type"`
				Value int `json:"value"`
			} `json:"officerOutput,omitempty"`
		} `json:"areas"`
		TroopLv            int           `json:"troopLv"`
		Luban              int           `json:"luban"`
		SlaveNum           int           `json:"slaveNum"`
		BuyGrabCost        int           `json:"buyGrabCost"`
		BuyGrabNum         int           `json:"buyGrabNum"`
		BuyGrabState       int           `json:"buyGrabState"`
		HuangchengSlaveNum int           `json:"huangchengSlaveNum"`
		CurrentStage       int           `json:"currentStage"`
		TotalStage         int           `json:"totalStage"`
		ReformStatus       int           `json:"reformStatus"`
		NextNum            int           `json:"nextNum"`
		Invest             []interface{} `json:"invest"`
		DisplayIron        int           `json:"displayIron"`
	}
}

func NewGetMainCityInfo() *GetMainCityInfo {
	p := &GetMainCityInfo{}
	p.Zh = "【建筑】获取主城信息"
	p.Cmd = cmd.BuildingGetMainCityInfo
	return p
}
func (c *GetMainCityInfo) Update1() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
func (c *GetMainCityInfo) Update2() {
	util.Map2Struct(c.Send.(string), &c.Send2)
	fmt.Println()
}
