package gcld

import (
	"sockets-proxy/gcld/cmd"
	"sockets-proxy/gcld/cmd/login"
	"sockets-proxy/util/log"
)

type Bot struct {
	userinfo *login.UserInfo
}
type M interface {
	Data() []byte
	Hex()
}
type B struct {
	action Action `json:"action"`
}
type Action struct {
	state int    `json:"state"`
	data  string `json:"data"`
}

func NewBot() Bot {
	return Bot{}
}

func (receiver Bot) Handle(msg interface{}) {
	switch msg.(type) {
	case sendData:
		receiver.sendData(msg.(sendData))
	case recData:
		receiver.recData(msg.(recData))
	default:
		panic("error.")
	}
}
func (receiver Bot) sendData(s sendData) {
	s.Print()
	switch s.Command {
	case cmd.LoginUser:
		if receiver.userinfo == nil {
			receiver.userinfo = login.NewUserInfo()
		}
		receiver.userinfo.UpdateSend(s.Body)
	case cmd.BuildingGetMainCity:
		log.Log.Println("发送获取建筑信息...")
	default:
		log.Log.Println("SEND:", s.Command)
	}
}
func (receiver Bot) recData(r recData) {
	r.Print()
	switch r.Command {
	case cmd.LoginUser:
		if receiver.userinfo == nil {
			receiver.userinfo = login.NewUserInfo()
		}
		receiver.userinfo.UpdateRec([]byte(r.Body))
	case cmd.BuildingGetMainCity:
		log.Log.Println("收到建筑信息...")
	case cmd.PushPlayer:
		log.Log.Println("收到广播信息...")
	case cmd.PushChat:
		log.Log.Println("收到聊天信息...")
	default:
		log.Log.Println("REC:", r.Command)
	}
}
