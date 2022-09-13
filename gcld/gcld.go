package gcld

import (
	"github.com/r3labs/sse/v2"
	"sockets-proxy/util/log"

	"time"
)

type Bot struct {
	SseServer     *sse.Server
	pluginManager *PluginManager
}

func NewBot() *Bot {
	bot := &Bot{}
	bot.pluginManager = NewPluginManager("./plugins")
	return bot
}

func (bot *Bot) Task() {
	for true {
		time.Sleep(time.Second * 5)
	}
}

func (bot *Bot) PushMsg(id, data string) {
	bot.SseServer.Publish(id, &sse.Event{Data: []byte(data)})
}
func (bot *Bot) Handle(msg Data) {
	if plugin, ok := bot.pluginManager.plugins[msg.Command]; ok {
		r, err := plugin.Vm.Call("Plugin.Receive", nil, msg.Type, msg.Body)
		if err != nil {
			log.Log.Errorf("call plugin %s Plugin.Receive error ->%s", plugin.Command, err.Error())
		}
		result := r.String()
		if result != "0" {
			log.Log.Infof(result)
		}

	} else {
		log.Log.Infof("\n【%s】\nCommand:%s\nBody:%s", msg.Type, msg.Command, msg.Body)
	}
}
