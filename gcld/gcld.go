package gcld

import (
	"github.com/r3labs/sse/v2"
	"sockets-proxy/gcld/cmd"
	"sockets-proxy/gcld/cmd/building"
	"sockets-proxy/gcld/cmd/login"
	"sockets-proxy/gcld/cmd/player"
	"sockets-proxy/gcld/cmd/push"
	"sockets-proxy/util/log"
	"time"
)

type Bot struct {
	Game
	SseServer *sse.Server
}

type Game struct {
	loginUserinfo        *login.UserInfo
	playerPlayerinfo     *player.GetPlayerInfo
	playerPlayerlist     *player.GetPlayerList
	buildingMaincityinfo *building.GetMainCityInfo
	pushPlayer           *push.Player
	pushBuilding         *push.Building
	pushNotice           *push.Notice
}

func NewBot() *Bot {
	bot := &Bot{}
	bot.loginUserinfo = login.NewUserInfo()
	bot.playerPlayerinfo = player.NewGetPlayerInfo()
	bot.playerPlayerlist = player.NewGetPlayerList()
	bot.buildingMaincityinfo = building.NewGetMainCityInfo()
	bot.pushPlayer = push.NewPlayer()
	bot.pushBuilding = push.NewBuilding()
	bot.pushNotice = push.NewNotice()
	go bot.Task()
	return bot
}

func (bot *Bot) Task() {
	for true {

		log.Log.Infof("当前登录用户：%s", bot.playerPlayerinfo.Rec2.Player.PlayerName)
		time.Sleep(time.Second * 5)
	}
}

func (bot *Bot) PushMsg(id, data string) {
	bot.SseServer.Publish(id, &sse.Event{Data: []byte(data)})
}
func (bot *Bot) Handle(msg interface{}) {
	switch msg.(type) {
	case cmd.SendData:
		bot.SendData(msg.(cmd.SendData))
	case cmd.RecData:
		bot.RecData(msg.(cmd.RecData))
	default:
		panic("error.")
	}
}

// SendData
//
//	@Description:
//	@receiver bot
//	@param s
func (bot *Bot) SendData(s cmd.SendData) {
	switch s.Command {
	case cmd.LoginUser:
		bot.loginUserinfo.UpdateSend(s.Body)
	case cmd.BuildingGetMainCityInfo:
		bot.buildingMaincityinfo.UpdateSend(s.Body)
	case cmd.PushPlayer:
		bot.pushPlayer.UpdateSend(s.Body)
	case cmd.PushBuilding:
		bot.pushBuilding.UpdateSend(s.Body)
	default:
		log.Log.Infof("SEND:%s\nBody:%s", s.Command, s.Body)
	}
}

// RecData
//
//	@Description:
//	@receiver bot
//	@param r
func (bot *Bot) RecData(r cmd.RecData) {
	switch r.Command {
	case cmd.LoginUser:
		bot.loginUserinfo.UpdateRec(r.BodyByte)
		bot.loginUserinfo.Update()
	case cmd.BuildingGetMainCityInfo:
		bot.buildingMaincityinfo.UpdateRec(r.BodyByte)
		bot.buildingMaincityinfo.Update()
	case cmd.PushPlayer:
		bot.pushPlayer.UpdateRec(r.BodyByte)
		bot.pushPlayer.Update()
	case cmd.PushBuilding:
		bot.pushBuilding.UpdateRec(r.BodyByte)
		bot.pushBuilding.Update()
	case cmd.PushNotice:
		bot.pushNotice.UpdateRec(r.BodyByte)
		bot.pushNotice.Update()
	default:
		log.Log.Infof("Rec:%s\nBody:%s", r.Command, r.Body)
	}
}
