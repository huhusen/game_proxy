package main

import (
	"embed"
	"encoding/hex"
	"fmt"
	"github.com/r3labs/sse/v2"
	"net/http"
	"os/exec"
	"sockets-proxy/gcld"
	"sockets-proxy/gcld/cmd/chariot"
	"sockets-proxy/proxy"
	"strconv"
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
		count := request.FormValue("count")
		cmd := request.FormValue("cmd")

		if cmd == "zc" {
			go func() {
				for _, proto := range proxyServer.Clients {
					r, ok := proto.(*proxy.Socket5)
					if ok {
						for t := 0; t < 100; t++ {
							for i := 1; i < 5; i++ {
								for j := 1; j < 5; j++ {
									chariot.NewForgeSpRequest(fmt.Sprintf("%d", i), fmt.Sprintf("%d", j)).Hex()
									r.Send2Server(chariot.NewForgeSpRequest(fmt.Sprintf("%d", i), fmt.Sprintf("%d", j)).Data())
								}
								chariot.NewGetBpInfoRequest(fmt.Sprintf("%d", i), "5").Hex()
								r.Send2Server(chariot.NewGetBpInfoRequest(fmt.Sprintf("%d", i), "5").Data())
								chariot.NewForgeBpInfoRequest(fmt.Sprintf("%d", i)).Hex()
								r.Send2Server(chariot.NewForgeBpInfoRequest(fmt.Sprintf("%d", i)).Data())
							}
						}

					}
				}
			}()

		} else {
			c, _ := strconv.Atoi(count)
			b, _ := hex.DecodeString(cmd)
			go func(count_ int, cmd_ []byte) {
				for i := 0; i < count_; i++ {
					fmt.Println(i)
					for _, proto := range proxyServer.Clients {
						r, ok := proto.(*proxy.Socket5)
						if ok {
							r.Send2Server(cmd_)
						}
					}

				}
			}(c, b)
		}

		writer.Write([]byte("send success."))
	})
	go func() {
		time.Sleep(time.Second * 1)
		cmd := exec.Command(`cmd`, `/c`, `start`, `http://127.0.0.1:36134/ui/`)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		//cmd.Start()
	}()
	http.ListenAndServe(":36134", nil)

	//n := pcap.NewCaptureUnPack("C:\\Users\\uma-pc001\\Desktop\\fsdownload\\gcld2.pcapng")
	//b := gcld.NewBot()
	//
	//n.DataHandle = func(data []byte, n pcap.Net) {
	//	//if strings.Index(hex.EncodeToString(data), "0d0a0d0a789c") != -1 {
	//	//spe := []byte{0x0d, 0x0a, 0x0d, 0x0a}
	//	//arr := bytes.Split(data, spe)
	//	//d := util.ZlibUnCompress(arr[1])
	//	//fmt.Println(string(d))
	//	//}
	//	if n.TCP.SrcPort == "9128" {
	//		b.Handle(gcld.NewRecvData(data))
	//	}
	//	if n.TCP.DstPort == "9128" {
	//		b.Handle(gcld.NewSendData(data))
	//	}
	//
	//}
	//n.Run()
}
