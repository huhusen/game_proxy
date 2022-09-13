package main

import (
	"embed"
	"github.com/r3labs/sse/v2"
	"net/http"
	"os/exec"
	"sockets-proxy/gcld"
	"sockets-proxy/proxy"
	"syscall"
	"time"
)

var (
	//go:embed ui
	res       embed.FS
	sseServer *sse.Server
	bot       *gcld.Bot
)

func main() {
	bot = gcld.NewBot()
	sseServer = sse.New()
	sseServer.CreateStream("msg")
	bot.SseServer = sseServer
	proxyServer := proxy.NewServer(1888)
	proxyServer.OnSocket5ResponseEvent = bot.OnSocket5ResponseEvent
	proxyServer.OnSocket5RequestEvent = bot.OnSocket5RequestEvent
	proxyServer.OnHttpRequestEvent = bot.OnHttpRequestEvent
	proxyServer.OnHttpResponseEvent = bot.OnHttpResponseEvent
	proxyServer.Start()
	http.Handle("/", http.FileServer(http.FS(res)))
	http.HandleFunc("/msg", sseServer.ServeHTTP)
	http.HandleFunc("/start", func(writer http.ResponseWriter, request *http.Request) {
		//go proxyServer.Start()
		writer.Write([]byte("proxyServer started."))
	})
	http.HandleFunc("/send", func(writer http.ResponseWriter, request *http.Request) {

		writer.Write([]byte("send success."))
	})
	go func() {
		time.Sleep(time.Second * 1)
		cmd := exec.Command(`cmd`, `/c`, `start`, `http://127.0.0.1:36134/ui/`)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		//cmd.Start()
	}()
	http.ListenAndServe(":36134", nil)
}
