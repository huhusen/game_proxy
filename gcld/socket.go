package gcld

import (
	"encoding/hex"
	"sockets-proxy/gcld/cmd"
)

func (bot *Bot) OnSocket5RequestEvent(message []byte) (out []byte) {
	s := cmd.NewSendData(message)
	bot.Handle(s)
	bot.PushMsg("msg", hex.EncodeToString(s.Data))
	bot.PushMsg("msg", s.String())
	return
}

func (bot *Bot) OnSocket5ResponseEvent(message []byte) (out []byte) {
	r := cmd.NewRecvData(message)
	bot.Handle(r)
	bot.PushMsg("msg", r.String())
	return
}
