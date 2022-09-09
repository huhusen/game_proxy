package gcld

import (
	"sockets-proxy/binaryext"
	"sockets-proxy/log"
	"sockets-proxy/util"
)

type Bot struct {
	userinfo *LoginUserInfo
}
type M interface {
	Data() []byte
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
func PacketData(cmd CMD, body string) []byte {
	bs := []byte(body)
	writer := binaryext.NewByteArray()
	writer.WriteInt(32 + 4 + len(bs))
	writer.Write([]byte(cmd))
	writer.WriteInt(0)
	writer.Write(util.NewBuf(32 - len(cmd)))
	writer.Write(bs)
	return writer.Data()
}
func (receiver Bot) Handle(msg interface{}) {
	switch msg.(type) {
	case sendData:

	case recData:

	default:
		panic("error.")
	}
}
func (receiver Bot) sendData(s sendData) {
	s.Print()
	switch s.Command {
	case LoginUser:
		if receiver.userinfo == nil {
			receiver.userinfo = NewLoginUserInfo()
		}
		receiver.userinfo.UpdateSend(s.Body)
	case BuildingGetMainCityInfo:
		log.Log.Println("发送获取建筑信息...")
	default:
		log.Log.Println("SEND:", s.Command)
	}
}
func (receiver Bot) recData(r recData) {
	switch r.Command {
	case LoginUser:
		if receiver.userinfo == nil {
			receiver.userinfo = NewLoginUserInfo()
		}
		receiver.userinfo.UpdateRec([]byte(r.Body))
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
