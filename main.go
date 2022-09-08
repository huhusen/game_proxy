package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sockets-proxy/gcld"
	"sockets-proxy/log"
	"sockets-proxy/pcap"
	"sockets-proxy/proxy"
	"sockets-proxy/util"
	"strings"
	"time"
)

func main() {

	n := pcap.NewCaptureUnPack("C:\\Users\\uma-pc001\\Desktop\\fsdownload\\gcld.pcapng")
	n.DataHandle = func(data []byte, n pcap.Net) {

		if strings.Index(hex.EncodeToString(data), "0d0a0d0a789c") != -1 {

			spe := []byte{0x0d, 0x0a, 0x0d, 0x0a}
			arr := bytes.Split(data, spe)
			d := util.ZlibUnCompress(arr[1])
			fmt.Println(string(d))
		}

		if n.TCP.SrcPort == "9128" {

			gcld.NewRecvData(data).Print()
		}
		if n.TCP.DstPort == "9128" {

			gcld.NewSendData(data).Print()
		}

	}
	n.Run()

	s := proxy.NewServer(1888)
	s.OnSocket5RequestEvent = func(message []byte) (out []byte) {
		log.Log.Println("send", string(message))
		if strings.Index(string(message), "123") != -1 {
			out = []byte("hack Send 123\n")
		}
		return
	}
	s.OnSocket5ResponseEvent = func(message []byte) (out []byte) {
		log.Log.Println("recv", string(message))
		if strings.Index(string(message), "123") != -1 {
			out = []byte("hack recv 123\n")
		}
		return
	}

	go s.Start()

	for {
		for _, proto := range s.Clients {
			r, ok := proto.(*proxy.Socket5)
			if ok {
				r.Send2Server([]byte("I am socket5 Client\n"))
				r.Send2client([]byte("I am socket5 Server\n"))
			}
		}
		time.Sleep(time.Second * 10)
	}
}
