package main

import (
	"embed"
	"encoding/hex"
	"fmt"
	"github.com/r3labs/sse/v2"
	"net/http"
	"os/exec"
	"sockets-proxy/gcld"
	"sockets-proxy/proxy"
	"strconv"
	"syscall"
	"time"
)

var (
	//go:embed ui
	res       embed.FS
	sseServer *sse.Server
	bot       gcld.Bot
)

func PushMsg(id, data string) {
	sseServer.Publish(id, &sse.Event{Data: []byte(data)})
}

func OnSocket5RequestEvent(message []byte) (out []byte) {
	s := gcld.NewSendData(message)
	bot.Handle(gcld.NewSendData(message))
	PushMsg("msg", hex.EncodeToString(s.Data))
	PushMsg("msg", s.String())
	return
}

func OnSocket5ResponseEvent(message []byte) (out []byte) {
	r := gcld.NewRecvData(message)
	bot.Handle(r)
	PushMsg("msg", r.String())
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
	go func() {
		time.Sleep(time.Second * 1)
		cmd := exec.Command(`cmd`, `/c`, `start`, `http://127.0.0.1:36134/ui/`)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Start()
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
