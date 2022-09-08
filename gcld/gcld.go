package gcld

import "sockets-proxy/log"

type Bot struct {
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
	case LoginUser:
		log.Log.Println("发送登录消息...")
	case BuildingGetMainCityInfo:
		log.Log.Println("发送获取建筑信息...")
	default:
		log.Log.Println("SEND:", s.Command)
	}
}
func (receiver Bot) recData(r recData) {
	r.Print()
	switch r.Command {
	case LoginUser:
		log.Log.Println("收到登录消息...")
	case BuildingGetMainCityInfo:
		log.Log.Println("收到建筑信息...")
	case PushPlayer:
		log.Log.Println("收到广播信息...")
	case PushChat:
		log.Log.Println("收到聊天信息...")
	default:
		log.Log.Println("REC:", r.Command)
	}
}
