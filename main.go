package main

import (
	"embed"
	"encoding/hex"
	"fmt"
	"github.com/r3labs/sse/v2"
	"net/http"
	"sockets-proxy/gcld"
	"sockets-proxy/proxy"
	"strconv"
	"time"
)

var (
	//go:embed ui
	res       embed.FS
	sseServer *sse.Server
	bot       gcld.Bot
)

func OnSocket5RequestEvent(message []byte) (out []byte) {
	sseServer.Publish("msg", &sse.Event{Data: []byte(hex.EncodeToString(message))})
	s := gcld.NewSendData(message)
	bot.Handle(gcld.NewSendData(message))
	sseServer.Publish("msg", &sse.Event{Data: s.Bytes()})
	return
}

func OnSocket5ResponseEvent(message []byte) (out []byte) {
	r := gcld.NewRecvData(message)
	bot.Handle(r)
	sseServer.Publish("msg", &sse.Event{Data: r.Bytes()})
	return
}

func main() {
	bot = gcld.NewBot()
	sseServer = sse.New()
	sseServer.CreateStream("msg")
	proxyServer := proxy.NewServer(1888)
	proxyServer.OnSocket5ResponseEvent = OnSocket5ResponseEvent
	proxyServer.OnSocket5RequestEvent = OnSocket5RequestEvent
	http.Handle("/", http.FileServer(http.FS(res)))
	http.HandleFunc("/msg", sseServer.ServeHTTP)
	http.HandleFunc("/start", func(writer http.ResponseWriter, request *http.Request) {
		go proxyServer.Start()
		writer.Write([]byte("proxyServer started."))
	})
	http.HandleFunc("/send", func(writer http.ResponseWriter, request *http.Request) {
		count := request.FormValue("count")
		cmd := request.FormValue("cmd")
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
				time.Sleep(time.Millisecond * 50)
			}
		}(c, b)
		writer.Write([]byte("send success."))
	})

	http.ListenAndServe(":8080", nil)

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
