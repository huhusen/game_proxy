package push

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

type Building struct {
	cmd.Command
	Rec2 struct {
		Output struct {
			OutputInfo []struct {
				BuildingType int `json:"buildingType"`
				Output       int `json:"output"`
				RealValue    int `json:"realValue"`
			} `json:"outputInfo"`
			IsChangeDisplaye bool `json:"isChangeDisplaye"`
		} `json:"output"`
	}
}

func NewBuilding() *Building {
	u := Building{}
	u.Cmd = cmd.PushBuilding
	u.Zh = "【推送】建筑信息"
	//u.Send = u.Send2
	u.Rec = u.Rec2
	return &u
}
func (c *Building) Update() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
