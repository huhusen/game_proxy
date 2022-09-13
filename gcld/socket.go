package gcld

import (
	"encoding/hex"
)

func (bot *Bot) OnSocket5RequestEvent(message []byte) (out []byte) {
	s := NewSendData(message)
	bot.Handle(s)
	bot.PushMsg("msg", hex.EncodeToString(s.Data))
	bot.PushMsg("msg", s.String())
	return
}

func (bot *Bot) OnSocket5ResponseEvent(message []byte) (out []byte) {
	r := NewRecvData(message)
	bot.Handle(r)
	bot.PushMsg("msg", r.String())
	return
}
