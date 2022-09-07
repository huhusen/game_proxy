package main

import (
	"sockets-proxy/log"
	"sockets-proxy/proxy"
	"strings"
)

func main() {

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

	s.Start()

	//for {
	//	for k, proto := range s.Clients {
	//		r, ok := proto.(*proxy.Socket5)
	//		if ok {
	//			r.Send2Server([]byte("I am socket5 Client"))
	//			r.Send2client([]byte("I am socket5 Server"))
	//		}
	//		fmt.Println(k)
	//
	//	}
	//	time.Sleep(time.Second * 10)
	//}
}
