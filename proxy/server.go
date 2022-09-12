package proxy

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"sockets-proxy/util/log"

	"time"
)

type Client struct {
	conn   net.Conn
	writer *bufio.Writer
	reader *bufio.Reader
	proxy  *Server
}
type Server struct {
	port                   int
	listener               *net.TCPListener
	OnSocket5ResponseEvent func(message []byte) (out []byte)
	OnSocket5RequestEvent  func(message []byte) (out []byte)
	OnHttpRequestEvent     func(request *http.Request)
	OnHttpResponseEvent    func(response *http.Response)
	Clients                map[string]interface{}
}
type Protocol interface {
	Handle()
}

func NewServer(port int) *Server {
	clients := make(map[string]interface{})
	return &Server{port: port, Clients: clients}
}
func (s *Server) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	s.listener = listener
	s.logo()
	s.multiListen()
	select {}
}
func (s *Server) logo() {
	logo := `
 _____  _____  __  __  _____       _____  _____  _____  __  __ ___ ___
/   __\/  _  \/  \/  \/   __\ ___ /  _  \/  _  \/  _  \/  \/  \\  |  /
|  |_ ||  _  ||  \/  ||   __|<___>|   __/|  _  <|  |  |>-    -< |   | 
\_____/\__|__/\__ \__/\_____/     \__/   \__|\_/\_____/\__/\__/ \___/ 
			game_proxy 1.0
`
	fmt.Println(logo)
	fmt.Println("代理监听端口:0.0.0.0:", s.port)
}
func (s *Server) multiListen() {
	for i := 0; i < 5; i++ {
		go func() {
			for {
				clientConn, err := s.listener.Accept()
				if err != nil {
					if e, ok := err.(net.Error); ok && e.Timeout() {
						log.Log.Println("accept connection failed ,reconnected：" + err.Error())
						time.Sleep(time.Second / 20)
					} else {
						log.Log.Println("accept connection failed: " + err.Error())
					}
					continue
				}
				go s.handle(clientConn)
			}
		}()
	}
}
func (s *Server) handle(conn net.Conn) {
	var protocol Protocol
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	peek, err := reader.Peek(1)
	if err != nil {
		return
	}
	peekHex := fmt.Sprintf("0x%x", peek[0])
	client := Client{proxy: s, conn: conn, writer: writer, reader: reader}
	switch peekHex {
	case "0x47", "0x43", "0x50", "0x4f", "0x44", "0x48":
		protocol = &ProxyHttp{client: client}
		break
	case "0x5":
		protocol = &Socket5{client: client, state: false}
		s.Clients = make(map[string]interface{})
		s.Clients[conn.RemoteAddr().String()] = protocol
	default:
		log.Log.Info("tcp.")
	}

	protocol.Handle()
}
