package log

import (
	_ "github.com/husanpao/logrus-easy-formatter"
	easy "github.com/husanpao/logrus-easy-formatter"
	"github.com/husanpao/timewriter"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetReportCaller(true)
	Log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% [%lvl%] : %msg%\n",
	})
	Log.SetLevel(logrus.InfoLevel)
	timeWriter := &timewriter.TimeWriter{
		Dir:           "./applog",
		Compress:      true,
		ReserveDay:    30,
		Screen:        true,
		LogFilePrefix: "game-proxy", // default is process name
	}
	Log.Out = timeWriter
}
